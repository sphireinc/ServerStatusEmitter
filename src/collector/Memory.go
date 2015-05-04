package collector

import (
	psutil_mem "github.com/shirou/gopsutil/mem"
)

type Memory struct {
	VirtualMemoryStat interface{} `json:"virtual_memory_stat"`
	SwapMemoryStat    interface{} `json:"swap_memory_stat"`
}

func (MemoryPtr *Memory) Collect() *Memory {
	MemoryPtr.VirtualMemoryStat, _ = psutil_mem.VirtualMemory()
	MemoryPtr.SwapMemoryStat, _ = psutil_mem.SwapMemory()

	return MemoryPtr
}
