package collector

import (
	psutil_mem "github.com/shirou/gopsutil/mem"
)

type Memory struct {
	VirtualMemoryStat VirtualMemoryStat `json:"virtual_memory_stat"`
	SwapMemoryStat    SwapMemoryStat    `json:"swap_memory_stat"`
}

type VirtualMemoryStat struct {
	Total       uint64  `json:"total"`
	Available   uint64  `json:"available"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"used_percent"`
	Free        uint64  `json:"free"`
	Active      uint64  `json:"active"`
	Inactive    uint64  `json:"inactive"`
	Buffers     uint64  `json:"buffers"`
	Cached      uint64  `json:"cached"`
	Wired       uint64  `json:"wired"`
	Shared      uint64  `json:"shared"`
}

type SwapMemoryStat struct {
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	Free        uint64  `json:"free"`
	UsedPercent float64 `json:"used_percent"`
	Sin         uint64  `json:"sin"`
	Sout        uint64  `json:"sout"`
}

func (MemoryPtr *Memory) Collect() *Memory {
	virtual, _ := psutil_mem.VirtualMemory()
	swap, _ := psutil_mem.SwapMemory()

	MemoryPtr.VirtualMemoryStat = VirtualMemoryStat{
		Total:       virtual.Total,
		Available:   virtual.Available,
		Used:        virtual.Used,
		UsedPercent: virtual.UsedPercent,
		Free:        virtual.Free,
		Active:      virtual.Active,
		Inactive:    virtual.Inactive,
		Buffers:     virtual.Buffers,
		Cached:      virtual.Cached,
		Wired:       virtual.Wired,
		Shared:      virtual.Shared,
	}

	MemoryPtr.SwapMemoryStat = SwapMemoryStat{
		Total:       swap.Total,
		Used:        swap.Used,
		Free:        swap.Free,
		UsedPercent: swap.UsedPercent,
		Sin:         swap.Sin,
		Sout:        swap.Sout,
	}

	return MemoryPtr
}
