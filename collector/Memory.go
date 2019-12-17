package main

import (
	"github.com/shirou/gopsutil/mem"
)

// Memory is the struct that contains data about the Memory
type Memory struct {
	VirtualMemoryStat interface{}
	SwapMemoryStat    interface{}
}

// Collect helps to collect data about the Memory
// and store it in the Memory struct
func (MemoryPtr *Memory) Collect() *Memory {
	MemoryPtr.VirtualMemoryStat, _ = mem.VirtualMemory()
	MemoryPtr.SwapMemoryStat, _ = mem.SwapMemory()

	return MemoryPtr
}
