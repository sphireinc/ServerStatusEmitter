#!/usr/bin/python

import sys
import json
import time
import sched
import logging
from lib import cpu, memory, disks, network, system, transport

__cache = []
__cache_timer = 60
__cache_keeper = 0

def main(scheduler, config, logger):
    global __cache
    global __cache_timer
    global __cache_keeper

    logger.info("Running scheduler")
    payload = {
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
    payload = json.dumps(payload)

    __cache.append(payload)

    if __cache_keeper < __cache_timer:
        __cache_keeper += config.get('interval')
        print __cache_keeper
    else:
        payload = {"payload": __cache}
        __cache_keeper = 0
        transport.Transport(payload, config, logger)
        __cache = []

    # Schedule a new run at the specified interval
    logger.info("Setting new scheduler")
    scheduler.enter(config.get('interval'), 1, main, (scheduler, config, logger))
    scheduler.run()

if __name__ == '__main__':
    try:
        config = (json.loads(open("config.json").read()))['mothership']

        log_level = logging.WARN if str(config.get('log_level')).upper() == "WARN" else logging.WARN
        logging.basicConfig(filename=config['log'], filemode='a',
                format='%(asctime)s,%(msecs)d %(name)s %(levelname)s %(message)s',
                datefmt='%H:%M:%S', level=log_level)
        logger = logging.getLogger('sse_logger')

        __cache_timer = config['cache_time'] if 'cache_time' in config else 60

        # If the interval is higher, just exit
        if config['interval'] > __cache_timer:
            print >> sys.stderr, "Report interval is higher than cache timer."
            sys.exit(1)

        scheduler = sched.scheduler(time.time, time.sleep)

        main(scheduler, config, logger)
    except KeyboardInterrupt:
        print >> sys.stderr, '\nExiting by user request.\n'
        sys.exit(0)
    except Exception as e:
        import traceback, os.path
        top = traceback.extract_stack()[-1]
        location = '\n' + type(e).__name__ + '@' + top[0]
        print >> sys.stderr, location, '=>', str(e)
        sys.exit(1)
