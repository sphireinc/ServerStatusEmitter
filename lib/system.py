import psutil

class System:
    def __init__(self):
        pass

    @staticmethod
    def users():
        return { "users": psutil.users() }

    @staticmethod
    def boot_time():
        return { "boot_time": psutil.boot_time() }
