package runner

import (
	"github.com/jsanc623/ServerStatusEmitter/collector"
	"github.com/jsanc623/ServerStatusEmitter/config"
	error2 "github.com/jsanc623/ServerStatusEmitter/sphlog"
	"time"
)

var Conf config.Config

// Snapshot struct is a collection of other structs
// which are relayed from the different segments of
// the collector package.
type Snapshot struct {
	CPU     *collector.CPU
	Disks   *collector.Disks
	Memory  *collector.Memory
	Network *collector.Network
	System  *collector.System
	Time    time.Time
}

// Collector collects a snapshot of the system at
// the time of calling and stores it in Snapshot struct.
func (Snapshot *Snapshot) Collector() {
	// Initialize collectors
	var CPU collector.CPU
	var Disks collector.Disks
	var Memory collector.Memory
	var Network collector.Network
	var System collector.System

	// Perform collection runs
	err := CPU.Collect()
	error2.LogError(err)

	err = Disks.Collect(Conf.Settings.Disk.IncludePartitionData)
	error2.LogError(err)

	err = Memory.Collect()
	error2.LogError(err)

	err = Network.Collect()
	error2.LogError(err)

	err = System.Collect(Conf.Settings.System.IncludeUsers)
	error2.LogError(err)

	Snapshot.Time = time.Now().UTC()
	Snapshot.CPU = &CPU
	Snapshot.Disks = &Disks
	Snapshot.Memory = &Memory
	Snapshot.Network = &Network
	Snapshot.System = &System
}
