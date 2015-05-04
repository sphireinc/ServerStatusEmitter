package collector

import (
	psutil_host "github.com/shirou/gopsutil/host"
	psutil_load "github.com/shirou/gopsutil/load"
)

type System struct {
	HostInfo interface{} `json:"host_info"`
	LoadAvg  interface{} `json:"load_avg"`
	Users    interface{} `json:"boot_time"`
}

func (SystemPtr *System) Collect() *System {
	SystemPtr.HostInfo, _ = psutil_host.HostInfo()
	SystemPtr.Users, _ = psutil_host.Users()
	SystemPtr.LoadAvg, _ = psutil_load.LoadAvg()

	return SystemPtr
}
