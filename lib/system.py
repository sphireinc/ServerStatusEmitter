class System:
    psutil = None

    def __init__(self, psutil):
        self.psutil = psutil

    def snapshot(self):
        return {
            "users": self.psutil.users(),
            "boot_time": self.psutil.boot_time()
        }