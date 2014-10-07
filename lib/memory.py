import psutil

class Memory:
    def __init__(self):
        pass

    @staticmethod
    def virtual_memory():
        return { "virtual": psutil.virtual_memory() }

    @staticmethod
    def swap_memory():
        return { "swap": psutil.swap_memory() }