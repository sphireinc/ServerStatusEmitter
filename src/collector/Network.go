package collector

import (
	psutil_net "github.com/shirou/gopsutil/net"
)

// Network is the struct that contains data about the Network
type Network struct {
	NetIOCounters interface{} `json:"io_counter"`
	NetInterface  interface{} `json:"interface"`
}

// Collect helps to collect data about the Network
// and store it in the Network struct
func (NetworkPtr *Network) Collect() *Network {
	NetworkPtr.NetIOCounters, _ = psutil_net.NetIOCounters(true)
	NetworkPtr.NetInterface, _ = psutil_net.NetInterfaces()

	return NetworkPtr
}
