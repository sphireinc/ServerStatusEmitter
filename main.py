#!/usr/bin/python

import sys
import json
import time
from lib import cpu, memory, disks, network, system, transport

def main():
    config = (json.loads(open("config.json").read()))['mothership']

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
    transport.Transport(payload, config)

if __name__ == '__main__':
    try:
        main()
    except KeyboardInterrupt:
        print >> sys.stderr, '\nExiting by user request.\n'
        sys.exit(1)
