#!/usr/bin/python

import sys
import json
import time
import sched
import socket
from lib import cpu, memory, disks, network, system, transport

__cache = []
__cache_timer = 0
__cache_keeper = 0

@profile
def main(scheduler, config, sock):
    global __cache
    global __cache_timer
    global __cache_keeper

    payload = {
        "_id": {
            "time": time.time(),
            "id": config['identification']['id'],
            "hostname": config['identification'].get('hostname', socket.gethostname()),
            "type": config['identification'].get('type', 'false')
        },
        "cpu": [
            cpu.CPU.cpu_count(),
            cpu.CPU.cpu_percent(),
            cpu.CPU.cpu_times()
        ],
        "memory": [
            memory.Memory.swap_memory(),
            memory.Memory.virtual_memory()
        ],
        "disks": [
            disks.Disks.disk_io_counters(),
            disks.Disks.disk_partitions(),
            disks.Disks.disk_usage()
        ],
        "network": [
            network.Network.net_connections(),
            network.Network.net_io_counters()
        ],
        "system": [
            system.System.users(),
            system.System.boot_time()
        ]
    }
    __cache.append(json.dumps(payload))

    if __cache_keeper < __cache_timer:
        __cache_keeper += config.get('interval')
    else:
        transport.Transport({"payload": __cache}, config, sock)
        __cache_keeper = 0
        __cache = []

    # Schedule a new run at the specified interval
    scheduler.enter(config['interval'], 1, main, (scheduler, config, sock))
    scheduler.run()

if __name__ == '__main__':
    try:
        config = (json.loads(open("config.json").read()))['config']

        config['disable_cache'] = False
        if config['cache'].get('enabled') is True:
            __cache_timer = config['cache'].get('time_seconds_to_cache_between_sends', 60)
            config['interval'] = config['cache'].get('interval_seconds_between_captures', 5)

            # If the interval is higher, just exit
            if config['interval'] > __cache_timer:
                print >> sys.stderr, "Report interval is higher than cache timer."
                sys.exit(1)

        sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
        scheduler = sched.scheduler(time.time, time.sleep)
        main(scheduler, config, sock)
    except KeyboardInterrupt:
        print >> sys.stderr, '\nExiting by user request.\n'
        sys.exit(0)
    except Exception as e:
        location = '\n' + type(e).__name__
        print >> sys.stderr, location, '=>', str(e)
        sys.exit(1)
