package collector

import (
	psutil_cpu "github.com/shirou/gopsutil/cpu"
)

type CPU struct {
	CPUTimesStat    interface{} `json:"cpu_time_stat"`
	CPUInfoStat     interface{} `json:"cpu_info_stat"`
	CPUCount        int         `json:"cpu_count"`
	CPUCountLogical int         `json:"cpu_count_logical"`
}

func (CPUPtr *CPU) Collect() *CPU {
	CPUPtr.CPUCountLogical, _ = psutil_cpu.CPUCounts(true)
	CPUPtr.CPUCount, _ = psutil_cpu.CPUCounts(false)
	CPUPtr.CPUTimesStat, _ = psutil_cpu.CPUTimes(true)
	CPUPtr.CPUInfoStat, _ = psutil_cpu.CPUInfo()

	return CPUPtr
}
