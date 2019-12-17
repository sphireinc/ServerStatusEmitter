package main

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
func (NetworkPtr *Network) Collect() *Network {
	NetworkPtr.NetIOCounters, _ = net.NetIOCounters(true)
	NetworkPtr.NetInterface, _ = net.NetInterfaces()

	return NetworkPtr
}
