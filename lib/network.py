import psutil

class Network:
    def __init__(self):
        pass

    @staticmethod
    def net_io_counters():
        return { "net_io_counters": psutil.net_io_counters() }

    @staticmethod
    def net_connections():
        return { "net_connections()": psutil.net_connections() }