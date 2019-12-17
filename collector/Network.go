package collector

import (
	"github.com/shirou/gopsutil/net"
)

// Network is the struct that contains data about the Network
type Network struct {
	NetIOCounters interface{}
	NetInterface  interface{}
}

// Collect helps to collect data about the Network
// and store it in the Network struct
func (Network *Network) Collect() error {
	var err error

	Network.NetIOCounters, err = net.IOCounters(true)
	if err != nil {
		return err
	}

	Network.NetInterface, err = net.Interfaces()
	if err != nil {
		return err
	}

	return nil
}
