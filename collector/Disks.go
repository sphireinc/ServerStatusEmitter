package collector

import (
	"github.com/shirou/gopsutil/disk"
)

// Disks is the struct that contains data about the Disks
type Disks struct {
	DiskUsage      interface{}
	DiskPartition  interface{}
	DiskIOCounters interface{}
}

// Collect helps to collect data about the Disks and store it in the Disks struct
func (Disks *Disks) Collect(diskPartition bool) error {
	var err error

	Disks.DiskUsage, err = disk.Usage("/")
	if err != nil {
		return err
	}

	Disks.DiskIOCounters, err = disk.IOCounters()
	if err != nil {
		return err
	}

	if diskPartition {
		Disks.DiskPartition, err = disk.Partitions(true)
		if err != nil {
			return err
		}
	}

	return nil
}
