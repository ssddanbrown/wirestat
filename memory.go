package main

import (
	"fmt"
	linuxproc "github.com/c9s/goprocinfo/linux"
)

type Memory struct {
	UsedPercent     uint   `json:"used_percent"`
	SwapUsedPercent uint   `json:"swap_used_percent"`
	Total           uint64 `json:"total"`
	Free            uint64 `json:"free"`
	Available       uint64 `json:"available"`
	Used            uint64 `json:"used"`
	SwapTotal       uint64 `json:"swap_total"`
	SwapFree        uint64 `json:"swap_free"`
	SwapUsed        uint64 `json:"swap_used"`
}

func GetMemory() (*Memory, error) {
	memInfo, err := linuxproc.ReadMemInfo("/proc/meminfo")
	if err != nil {
		return nil, fmt.Errorf("failed to stat meminfo data, received error: %s", err.Error())
	}

	used := memInfo.MemTotal - memInfo.MemFree
	swapUsed := memInfo.SwapTotal - memInfo.SwapFree

	mem := &Memory{
		UsedPercent:     uint((float64(used) / float64(memInfo.MemTotal)) * 100),
		SwapUsedPercent: uint((float64(swapUsed) / float64(memInfo.SwapTotal)) * 100),
		Total:           kibiToMega(memInfo.MemTotal),
		Free:            kibiToMega(memInfo.MemFree),
		Available:       kibiToMega(memInfo.MemAvailable),
		Used:            kibiToMega(used),
		SwapTotal:       kibiToMega(memInfo.SwapTotal),
		SwapFree:        kibiToMega(memInfo.SwapFree),
		SwapUsed:        kibiToMega(swapUsed),
	}

	return mem, nil
}

func kibiToMega(kibiVal uint64) uint64 {
	return uint64(float64(kibiVal) / 976.5625)
}
