#!/usr/bin/python
from hashlib import md5
import os
import sys
import json
import time
import sched
import socket

import psutil
import requests

from lib import cpu, memory, disks, network, system, transport

_cache = []
_cache_timer = 0
_cache_keeper = 0
_version = 1.0

def main(scheduler, config, sock, hostname, callers):
    global _cache
    global _cache_timer
    global _cache_keeper
    global _version

    payload = {
        "_id": {
            "time": time.time(),
            "id": config['identification']['id'],
            "hostname": hostname,
            "type": config['identification']['type']
        },
        "cpu": callers['cpu'].snapshot(),
        "memory": callers['memory'].snapshot(),
        "disks": callers['disks'].snapshot(),
        "network": callers['network'].snapshot(),
        "system": callers['system'].snapshot(),
        "version": _version
    }
    payload["digest"] = md5(str(payload)).hexdigest()

    _cache.append(payload)
    del payload

    if _cache_keeper < _cache_timer:
        _cache_keeper += config['interval']
    else:
        transport.Transport({ "payload": json.dumps(_cache) }, config, sock)
        _cache_keeper = 0
        _cache = []

    # Schedule a new run at the specified interval
    scheduler.enter(config['interval'], 1, main, (scheduler, config, sock, hostname, callers))
    scheduler.run()


def register(config):
    """
    Register this server/device with the mothership
    """
    return True


if __name__ == '__main__':
    try:
        config = (json.loads(open(os.path.dirname(os.path.abspath(__file__)) + "/config.json").read()))['config']
        config['identification']['type'] = config['identification'].get('type', 'false')

        config['disable_cache'] = False
        if config['cache'].get('enabled') is True:
            _cache_timer = config['cache'].get('time_seconds_to_cache_between_sends', 60)
            config['interval'] = config['cache'].get('interval_seconds_between_captures', 5)

            # If the interval is higher, just exit
            if config['interval'] > _cache_timer:
                print >> sys.stderr, "Report interval is higher than cache timer."
                sys.exit(1)

        sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
        scheduler = sched.scheduler(time.time, time.sleep)
        hostname = config['identification'].get('hostname', socket.gethostname())
        callers = {
            "cpu": cpu.CPU(psutil),
            "memory": memory.Memory(psutil),
            "disks": disks.Disks(psutil),
            "network": network.Network(psutil),
            "system": system.System(psutil)
        }
        if register(3):
            main(scheduler, config, sock, hostname, callers)
    except KeyboardInterrupt:
        print >> sys.stderr, '\nExiting by user request.\n'
        sys.exit(0)
    except Exception as e:
        location = '\n' + type(e).__name__
        print >> sys.stderr, location, '=>', str(e)
        sys.exit(1)
