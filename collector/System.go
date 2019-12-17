package main

import (
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
)

// System is the struct that contains data about the System
type System struct {
	HostInfo interface{}
	LoadAvg  interface{}
	Users    interface{}
}

// Collect helps to collect data about the System
// and store it in the System struct
func (SystemPtr *System) Collect(users bool) *System {
	SystemPtr.HostInfo, _ = host.HostInfo()
	SystemPtr.LoadAvg, _ = load.LoadAvg()

	if users {
		SystemPtr.Users, _ = host.Users()
	}

	return SystemPtr
}
