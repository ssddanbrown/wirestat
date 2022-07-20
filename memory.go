package main

import (
	"fmt"
	linuxproc "github.com/c9s/goprocinfo/linux"
)

func GetMemory() (map[string]uint64, error) {
	memInfo, err := linuxproc.ReadMemInfo("/proc/meminfo")
	if err != nil {
		return nil, fmt.Errorf("failed to stat meminfo data, received error: %s", err.Error())
	}

	used := memInfo.MemTotal - memInfo.MemFree
	swapUsed := memInfo.SwapTotal - memInfo.SwapFree

	mem := map[string]uint64{
		"used_percent":      uint64((float64(used) / float64(memInfo.MemTotal)) * 100),
		"swap_used_percent": uint64((float64(swapUsed) / float64(memInfo.SwapTotal)) * 100),
		"total":             kibiToMega(memInfo.MemTotal),
		"free":              kibiToMega(memInfo.MemFree),
		"available":         kibiToMega(memInfo.MemAvailable),
		"used":              kibiToMega(used),
		"swap_total":        kibiToMega(memInfo.SwapTotal),
		"swap_free":         kibiToMega(memInfo.SwapFree),
		"swap_used":         kibiToMega(swapUsed),
	}

	return mem, nil
}

func kibiToMega(kibiVal uint64) uint64 {
	return uint64(float64(kibiVal) / 976.5625)
}
