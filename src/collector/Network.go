package collector

import (
	psutil_net "github.com/shirou/gopsutil/net"
)

type Network struct {
	NetIOCounters interface{} `json:"io_counter"`
	NetInterface  interface{} `json:"interface"`
}

func (NetworkPtr *Network) Collect() *Network {
	NetworkPtr.NetIOCounters, _ = psutil_net.NetIOCounters(true)
	NetworkPtr.NetInterface, _ = psutil_net.NetInterfaces()

	return NetworkPtr
}
