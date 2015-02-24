class Memory:
    psutil = None

    def __init__(self, psutil):
        self.psutil = psutil

    def snapshot(self):
        """
        Generate a snapshot of the current memory state
        """
        return {
            "virtual": self.psutil.virtual_memory(),
            "swap": self.psutil.swap_memory()
        }