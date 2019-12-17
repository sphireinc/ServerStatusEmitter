package main

import (
	"github.com/shirou/gopsutil/cpu"
)

// CPU is the struct that contains data about the CPU
type CPU struct {
	Times        interface{}
	Info         interface{}
	Count        int
	CountLogical int
}

// Collect helps to collect data about the CPU
// and store it in the CPU struct
func (CPUPtr *CPU) Collect() *CPU {
	CPUPtr.CountLogical, _ = cpu.CPUCounts(true)
	CPUPtr.Count, _ = cpu.CPUCounts(false)
	CPUPtr.Times, _ = cpu.CPUTimes(true)
	CPUPtr.Info, _ = cpu.CPUInfo()

	return CPUPtr
}
