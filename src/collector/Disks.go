package collector

import (
	psutil_disk "github.com/shirou/gopsutil/disk"
)


// Disks is the struct that contains data about the Disks
type Disks struct {
	DiskUsage      interface{} `json:"disk_usage_stat"`
	DiskPartition  interface{} `json:"disk_partition_stat"`
	DiskIOCounters interface{} `json:"disk_io_counters_stat"`
}

// Collect helps to collect data about the Disks
// and store it in the Disks struct
func (DisksPtr *Disks) Collect(disk_partition bool) *Disks {
	DisksPtr.DiskUsage, _ = psutil_disk.DiskUsage("/")
	DisksPtr.DiskIOCounters, _ = psutil_disk.DiskIOCounters()

	if disk_partition {
		DisksPtr.DiskPartition, _ = psutil_disk.DiskPartitions(true)
	}

	return DisksPtr
}
