class CPU:
    psutil = None

    def __init__(self, psutil):
        self.psutil = psutil

    def snapshot(self):
        """
        Generate a snapshot of the current CPU state
        """
        cpu_time = self.psutil.cpu_times()

        return {
            "cpu_percent": self.psutil.cpu_percent(interval=1, percpu=True),
            "cpu_times": {
                "user": cpu_time[0],
                "system": cpu_time[1],
                "idle": cpu_time[2]
            },
            "cpu_count": {
                "virtual": self.psutil.cpu_count(),
                "physical": self.psutil.cpu_count(logical=False)
            },
            "load_average": {

            }
        }
