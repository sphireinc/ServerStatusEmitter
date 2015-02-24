class Disks:
    psutil = None

    def __init__(self, psutil):
        self.psutil = psutil

    def snapshot(self):
        return {
            "disk_usage": self.psutil.disk_usage('/'),
            "disk_partitions": self.psutil.disk_partitions(),
            "disk_io_counters": self.psutil.disk_io_counters()
        }
