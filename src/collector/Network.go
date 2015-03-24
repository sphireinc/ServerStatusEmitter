package collector

type Network struct {
	Addr              Addr              `json:"address"`
	NetInterfaceAddr  NetInterfaceAddr  `json:"network_interface_address"`
	NetIOCountersStat NetIOCountersStat `json:"network_io_counter_stat"`
	NetConnectionStat NetConnectionStat `json:"network_connection_stat"`
	NetInterfaceStat  NetInterfaceStat  `json:"network_interface_stat"`
}

type Addr struct {
	IP   string `json:"ip"`
	Port uint32 `json:"port"`
}

type NetInterfaceAddr struct {
	Addr string `json:"addr"`
}

type NetIOCountersStat struct {
	Name        string `json:"name"`         // interface name
	BytesSent   uint64 `json:"bytes_sent"`   // number of bytes sent
	BytesRecv   uint64 `json:"bytes_recv"`   // number of bytes received
	PacketsSent uint64 `json:"packets_sent"` // number of packets sent
	PacketsRecv uint64 `json:"packets_recv"` // number of packets received
	Errin       uint64 `json:"errin"`        // total number of errors while receiving
	Errout      uint64 `json:"errout"`       // total number of errors while sending
	Dropin      uint64 `json:"dropin"`       // total number of incoming packets which were dropped
	Dropout     uint64 `json:"dropout"`      // total number of outgoing packets which were dropped (always 0 on OSX and BSD)
}

type NetConnectionStat struct {
	Fd     uint32 `json:"fd"`
	Family uint32 `json:"family"`
	Type   uint32 `json:"type"`
	Laddr  Addr   `json:"localaddr"`
	Raddr  Addr   `json:"remoteaddr"`
	Status string `json:"status"`
	Pid    int32  `json:"pid"`
}

type NetInterfaceStat struct {
	MTU          int                `json:"mtu"`          // maximum transmission unit
	Name         string             `json:"name"`         // e.g., "en0", "lo0", "eth0.100"
	HardwareAddr string             `json:"hardwareaddr"` // IEEE MAC-48, EUI-48 and EUI-64 form
	Flags        []string           `json:"flags"`        // e.g., FlagUp, FlagLoopback, FlagMulticast
	Addrs        []NetInterfaceAddr `json:"addrs"`
}

func (NetworkPtr *Network) Collect() *Network {

	return NetworkPtr
}
