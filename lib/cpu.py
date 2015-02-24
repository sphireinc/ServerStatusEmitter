from subprocess import check_output

class CPU:
    psutil = None
    cpu_count = {}
    cpu_passthrough = 0

    def __init__(self, psutil):
        self.psutil = psutil

    def snapshot(self):
        """
        Generate a snapshot of the current CPU state
        """
        cpu_time = self.psutil.cpu_times()

        # Only update the CPU counts every 100th pass through
        if self.cpu_count == {} or self.cpu_passthrough % 100 == 0:
            self.cpu_count = {
                "virtual": self.psutil.cpu_count(),
                "physical": self.psutil.cpu_count(logical=False)
            }

        print check_output(["cat", "/proc/loadavg"])

        return {
            "cpu_percent": self.psutil.cpu_percent(interval=1, percpu=True),
            "cpu_times": {
                "user": cpu_time[0],
                "system": cpu_time[1],
                "idle": cpu_time[2]
            },
            "cpu_count": self.cpu_count,
            "load_average": check_output(["cat", "/proc/loadavg"])
        }
