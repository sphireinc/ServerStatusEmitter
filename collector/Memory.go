package collector

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
func (Memory *Memory) Collect() error {
	var err error

	Memory.VirtualMemoryStat, err = mem.VirtualMemory()
	if err != nil {
		return err
	}

	Memory.SwapMemoryStat, err = mem.SwapMemory()
	if err != nil {
		return err
	}

	return nil
}
