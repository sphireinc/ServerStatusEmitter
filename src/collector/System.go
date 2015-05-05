package collector

import (
	psutil_host "github.com/shirou/gopsutil/host"
	psutil_load "github.com/shirou/gopsutil/load"
)

type System struct {
	HostInfo interface{} `json:"host_info"`
	LoadAvg  interface{} `json:"load_avg"`
	Users    interface{} `json:"users"`
}

func (SystemPtr *System) Collect(users bool) *System {
	SystemPtr.HostInfo, _ = psutil_host.HostInfo()
	SystemPtr.LoadAvg, _ = psutil_load.LoadAvg()

	if users {
		SystemPtr.Users, _ = psutil_host.Users()
	}

	return SystemPtr
}
