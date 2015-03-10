class System:
    __slots__ = ['psutil']

    psutil = None

    def __init__(self, psutil):
        self.psutil = psutil

    def snapshot(self):
        """
        Generate a snapshot of the current system state
        """
        return {
            "users": self.psutil.users(),
            "boot_time": self.psutil.boot_time()
        }
