import psutil

class Disks:
    def __init__(self):
        pass

    @staticmethod
    def disk_partitions():
        return { "disk_partitions": psutil.disk_partitions() }

    @staticmethod
    def disk_usage():
        return { "disk_usage": psutil.disk_usage('/') }

    @staticmethod
    def disk_io_counters():
        return { "disk_io_counters": psutil.disk_io_counters() }
