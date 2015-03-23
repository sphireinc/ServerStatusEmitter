package collector

import (
	psutil_cpu "github.com/shirou/gopsutil/cpu"
)

// 4:41p

type CPU struct {
	CPUTimesStat    CPUTimesStat `json:"cpu_time_stat"`
	CPUInfoStat     CPUInfoStat  `json:"cpu_info_stat"`
	CPUCount        int          `json:"cpu_count"`
	CPUCountLogical int          `json:"cpu_count_logical"`
}

type CPUTimesStat struct {
	CPU       string  `json:"cpu"`
	User      float64 `json:"user"`
	System    float64 `json:"system"`
	Idle      float64 `json:"idle"`
	Nice      float64 `json:"nice"`
	Iowait    float64 `json:"iowait"`
	Irq       float64 `json:"irq"`
	Softirq   float64 `json:"softirq"`
	Steal     float64 `json:"steal"`
	Guest     float64 `json:"guest"`
	GuestNice float64 `json:"guest_nice"`
	Stolen    float64 `json:"stolen"`
}

type CPUInfoStat struct {
	CPU        int32    `json:"cpu"`
	VendorID   string   `json:"vendor_id"`
	Family     string   `json:"family"`
	Model      string   `json:"model"`
	Stepping   int32    `json:"stepping"`
	PhysicalID string   `json:"physical_id"`
	CoreID     string   `json:"core_id"`
	Cores      int32    `json:"cores"`
	ModelName  string   `json:"model_name"`
	Mhz        float64  `json:"mhz"`
	CacheSize  int32    `json:"cache_size"`
	Flags      []string `json:"flags"`
}

func (CPUPtr *CPU) Collect() <-chan *CPU {
	out := make(chan *CPU)

	go func() {
		//cpu_time, _ := psutil_cpu.CPUTimes(true)
		//cpu_info, _ := psutil_cpu.CPUInfo()

		CPUPtr.CPUCountLogical, _ = psutil_cpu.CPUCounts(true)
		CPUPtr.CPUCount, _ = psutil_cpu.CPUCounts(false)

		/*CPU.CPUTimesStat = CPUTimesStat{
			CPU: cpu_time.CPU,
			User: cpu_time.User,
			System: cpu_time.System,
			Idle: cpu_time.Idle,
			Nice: cpu_time.Nice,
			Iowait: cpu_time.Iowait,
			Irq: cpu_time.Irq,
			Softirq: cpu_time.Softirq,
			Steal: cpu_time.Steal,
			Guest: cpu_time.Guest,
			GuestNice: cpu_time.GuestNice,
			Stolen: cpu_time.Stolen,
		}

		CPU.CPUInfoStat = CPUInfoStat{
			CPU: cpu_info.CPU,
			VendorID: cpu_info.VendorID,
			Family: cpu_info.Family,
			Model: cpu_info.Model,
			Stepping: cpu_info.Stepping,
			PhysicalID: cpu_info.PhysicalID,
			CoreID: cpu_info.CoreID,
			Cores: cpu_info.Cores,
			ModelName: cpu_info.ModelName,
			Mhz: cpu_info.Mhz,
			CacheSize: cpu_info.CacheSize,
			Flags: cpu_info.Flags,
		}*/

		out <- CPUPtr
		close(out)
	}()

	return out
}
