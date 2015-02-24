class Network:
    psutil = None

    def __init__(self, psutil):
        self.psutil = psutil

    def snapshot(self):
        """
        Generate a snapshot of the current network state
        """
        return {
            "net_io_counters": self.psutil.net_io_counters(),
            #"net_connections()": self.psutil.net_connections()
        }