class Disks:
    __slots__ = ['psutil', 'disk_usage', 'disk_partitions', 'disk_passthrough']

    psutil = None
    disk_usage = None
    disk_partitions = None
    disk_passthrough = 0

    def __init__(self, psutil):
        self.psutil = psutil

    def snapshot(self):
        """
        Generate a snapshot of the current disk state
        """
        # Only grab the disk partitions every 25th pass
        if self.disk_partitions is None or self.disk_passthrough % 25 == 0:
            self.disk_partitions = self.psutil.disk_partitions(all=True)

        # Only grab the disk usage every 5th pass
        if self.disk_usage is None or self.disk_passthrough % 5 == 0:
            self.disk_usage = self.psutil.disk_usage('/')

        self.disk_passthrough += 1

        return {
            "disk_usage": self.disk_usage,
            "disk_partitions": self.disk_partitions,
            "disk_io_counters": self.psutil.disk_io_counters(perdisk=True)
        }
