package main

import (
	"fmt"
	linuxproc "github.com/c9s/goprocinfo/linux"
)

type Memory struct {
	usedPercent     uint
	swapUsedPercent uint
	total           uint64
	free            uint64
	available       uint64
	used            uint64
	swapTotal       uint64
	swapFree        uint64
	swapUsed        uint64
}

func GetMemory() (*Memory, error) {
	memInfo, err := linuxproc.ReadMemInfo("/proc/meminfo")
	if err != nil {
		return nil, fmt.Errorf("failed to stat meminfo data, received error: %s", err.Error())
	}

	used := memInfo.MemTotal - memInfo.MemFree
	swapUsed := memInfo.SwapTotal - memInfo.SwapFree

	mem := &Memory{
		usedPercent:     uint(float64(used/memInfo.MemTotal) * 100),
		swapUsedPercent: uint(float64(swapUsed/memInfo.SwapTotal) * 100),
		total:           kibiToMega(memInfo.MemTotal),
		free:            kibiToMega(memInfo.MemFree),
		available:       kibiToMega(memInfo.MemAvailable),
		used:            kibiToMega(used),
		swapTotal:       kibiToMega(memInfo.SwapTotal),
		swapFree:        kibiToMega(memInfo.SwapFree),
		swapUsed:        kibiToMega(swapUsed),
	}

	return mem, nil
}

func kibiToMega(kibiVal uint64) uint64 {
	return uint64(float64(kibiVal) / 976.5625)
}
