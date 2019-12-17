package collector

import (
	"github.com/shirou/gopsutil/cpu"
)

// CPU is the struct that contains data about the CPU
type CPU struct {
	Times        []cpu.TimesStat
	Info         []cpu.InfoStat
	Count        int
	CountLogical int
}

// Collect helps to collect data about the CPU and store it in the CPU struct
func (CPU *CPU) Collect() error {
	var err error

	CPU.CountLogical, err = cpu.Counts(true)
	if err != nil {
		return err
	}

	CPU.Count, err = cpu.Counts(false)
	if err != nil {
		return err
	}

	CPU.Times, err = cpu.Times(true)
	if err != nil {
		return err
	}

	CPU.Info, err = cpu.Info()
	if err != nil {
		return err
	}

	return nil
}
