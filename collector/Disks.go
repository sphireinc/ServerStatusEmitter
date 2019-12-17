package main

import (
	"github.com/shirou/gopsutil/disk"
)

// Disks is the struct that contains data about the Disks
type Disks struct {
	DiskUsage      interface{}
	DiskPartition  interface{}
	DiskIOCounters interface{}
}

// Collect helps to collect data about the Disks
// and store it in the Disks struct
func (DisksPtr *Disks) Collect(diskPartition bool) *Disks {
	DisksPtr.DiskUsage, _ = disk.DiskUsage("/")
	DisksPtr.DiskIOCounters, _ = disk.DiskIOCounters()

	if diskPartition {
		DisksPtr.DiskPartition, _ = disk.DiskPartitions(true)
	}

	return DisksPtr
}
