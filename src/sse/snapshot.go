package sse

import (
	"collector"
	"time"
)

var (
	// CPU collector
	CPU        = collector.CPU{}

	// Disks collector
	Disks      = collector.Disks{}

	// Memory collector
	Memory    = collector.Memory{}

	// Network collector
	Network  = collector.Network{}

	// System collector
	System   = collector.System{}
)

// Snapshot struct is a collection of other structs
// which are relayed from the different segments of
// the collector package.
type Snapshot struct {
	CPU     *collector.CPU     `json:"cpu"`
	Disks   *collector.Disks   `json:"disks"`
	Memory  *collector.Memory  `json:"memory"`
	Network *collector.Network `json:"network"`
	System  *collector.System  `json:"system"`
	Time    time.Time          `json:"system_time"`
}

// Collector collects a snapshot of the system at
// the time of calling and stores it in Snapshot struct.
func (Snapshot *Snapshot) Collector(IncludePartitionData bool, IncludeUsers bool) *Snapshot {
	Snapshot.Time = time.Now().Local()
	Snapshot.CPU = CPU.Collect()
	Snapshot.Disks = Disks.Collect(IncludePartitionData)
	Snapshot.Memory = Memory.Collect()
	Snapshot.Network = Network.Collect()
	Snapshot.System = System.Collect(IncludeUsers)

	return Snapshot
}
