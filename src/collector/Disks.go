package collector

import (
	psutil_disk "github.com/shirou/gopsutil/disk"
)

type Disks struct {
	DiskUsage      interface{} `json:"disk_usage_stat"`
	DiskPartition  interface{} `json:"disk_partition_stat"`
	DiskIOCounters interface{} `json:"disk_io_counters_stat"`
}

func (DisksPtr *Disks) Collect() *Disks {
	DisksPtr.DiskUsage, _ = psutil_disk.DiskUsage("/")
	DisksPtr.DiskPartition, _ = psutil_disk.DiskPartitions(true)
	DisksPtr.DiskIOCounters, _ = psutil_disk.DiskIOCounters()

	return DisksPtr
}
