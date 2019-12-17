package collector

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

// Collect helps to collect data about the System and store it in the System struct
func (SystemPtr *System) Collect(users bool) error {
	var err error

	SystemPtr.HostInfo, err = host.Info()
	if err != nil {
		return err
	}

	SystemPtr.LoadAvg, err = load.Avg()
	if err != nil {
		return err
	}

	if users {
		SystemPtr.Users, err = host.Users()
		if err != nil {
			return err
		}
	}

	return nil
}
