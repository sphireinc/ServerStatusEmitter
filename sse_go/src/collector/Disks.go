package collector

type Disks struct {
	DiskUsageStat      DiskUsageStat
	DiskPartitionStat  DiskPartitionStat
	DiskIOCountersStat DiskIOCountersStat
}

type DiskUsageStat struct {
	Path              string  `json:"path"`
	Total             uint64  `json:"total"`
	Free              uint64  `json:"free"`
	Used              uint64  `json:"used"`
	UsedPercent       float64 `json:"used_percent"`
	InodesTotal       uint64  `json:"inodes_total"`
	InodesUsed        uint64  `json:"inodes_used"`
	InodesFree        uint64  `json:"inodes_free"`
	InodesUsedPercent float64 `json:"inodes_used_percent"`
}

type DiskPartitionStat struct {
	Device     string `json:"device"`
	Mountpoint string `json:"mountpoint"`
	Fstype     string `json:"fstype"`
	Opts       string `json:"opts"`
}

type DiskIOCountersStat struct {
	ReadCount    uint64 `json:"read_count"`
	WriteCount   uint64 `json:"write_count"`
	ReadBytes    uint64 `json:"read_bytes"`
	WriteBytes   uint64 `json:"write_bytes"`
	ReadTime     uint64 `json:"read_time"`
	WriteTime    uint64 `json:"write_time"`
	Name         string `json:"name"`
	IoTime       uint64 `json:"io_time"`
	SerialNumber string `json:"serial_number"`
}

func (Disks *Disks) Collect() *Disks {

	return Disks
}
