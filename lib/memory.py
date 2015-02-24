class Memory:
    psutil = None

    def __init__(self, psutil):
        self.psutil = psutil

    def snapshot(self):
        """
        Generate a snapshot of the current memory state
        """
        virtual_memory = self.psutil.virtual_memory
        swap_memory = self.psutil.swap_memory()

        return {
            "virtual": virtual_memory._asdict(),
            "swap": swap_memory._asdict()
        }