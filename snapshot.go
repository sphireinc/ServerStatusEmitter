package main

import (
	"time"
)

// CPU collector
var CPU CPU = CPU{}

// Disks collector
var Disks = Disks{}

// Memory collector
var Memory = Memory{}

// Network collector
var Network = Network{}

// System collector
var System = System{}

// Snapshot struct is a collection of other structs
// which are relayed from the different segments of
// the collector package.
type Snapshot struct {
	CPU     *CPU
	Disks   *Disks
	Memory  *Memory
	Network *Network
	System  *System
	Time    time.Time
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
