package collector

import (
	psutil_cpu "github.com/shirou/gopsutil/cpu"
)

type CPU struct {
	Times        interface{} `json:"cpu_time_stat"`
	Info         interface{} `json:"cpu_info_stat"`
	Count        int         `json:"cpu_count"`
	CountLogical int         `json:"cpu_count_logical"`
}

func (CPUPtr *CPU) Collect() *CPU {
	CPUPtr.CountLogical, _ = psutil_cpu.CPUCounts(true)
	CPUPtr.Count, _ = psutil_cpu.CPUCounts(false)
	CPUPtr.Times, _ = psutil_cpu.CPUTimes(true)
	CPUPtr.Info, _ = psutil_cpu.CPUInfo()

	return CPUPtr
}
