import psutil

class CPU:
    def __init__(self):
        pass

    @staticmethod
    def cpu_times():
        """
        Return system CPU times as a namedtuple
        """
        cpu_time = psutil.cpu_times()

        return {
            "cpu_times": {
                "user": cpu_time[0],
                "system": cpu_time[1],
                "idle": cpu_time[2]
            }
        }

    @staticmethod
    def cpu_percent():
        """
        Return a list of floats representing the current system-wide CPU utilization as a percentage.
        """
        return { "cpu_percent": psutil.cpu_percent(interval=1, percpu=True) }


    @staticmethod
    def cpu_count():
        """
        Return the number of logical CPUs in the system
        """
        return { "cpu_count": { "virtual": psutil.cpu_count(), "physical": psutil.cpu_count(logical=False) } }

    @staticmethod
    def cpu_load():
        """
        Return the load average
        """
        return { "load_average": { } }
