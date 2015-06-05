package collector

import (
	psutil_mem "github.com/shirou/gopsutil/mem"
)

// Memory is the struct that contains data about the Memory
type Memory struct {
	VirtualMemoryStat interface{} `json:"virtual_memory_stat"`
	SwapMemoryStat    interface{} `json:"swap_memory_stat"`
}

// Collect helps to collect data about the Memory
// and store it in the Memory struct
func (MemoryPtr *Memory) Collect() *Memory {
	MemoryPtr.VirtualMemoryStat, _ = psutil_mem.VirtualMemory()
	MemoryPtr.SwapMemoryStat, _ = psutil_mem.SwapMemory()

	return MemoryPtr
}
