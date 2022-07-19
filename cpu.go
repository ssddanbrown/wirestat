package main

import (
	"fmt"
	linuxproc "github.com/c9s/goprocinfo/linux"
	"time"
)

type CpuMap map[string]uint

func GetCpuMap() (CpuMap, error) {
	statA, err := linuxproc.ReadStat("/proc/stat")
	statErrMsg := "failed to stat CPU data, received error: %s"
	if err != nil {
		return nil, fmt.Errorf(statErrMsg, err.Error())
	}

	time.Sleep(time.Second)

	statB, err := linuxproc.ReadStat("/proc/stat")
	if err != nil {
		return nil, fmt.Errorf(statErrMsg, err.Error())
	}

	resultMap := make(map[string]uint)
	resultMap["all"] = cpuStatToPercent(statA.CPUStatAll, statB.CPUStatAll)
	for idx, statB := range statB.CPUStats {
		statA := statA.CPUStats[idx]
		resultMap[statB.Id] = cpuStatToPercent(statA, statB)
	}

	return resultMap, nil
}

func cpuStatToPercent(statA, statB linuxproc.CPUStat) uint {
	aIdle := statA.Idle + statA.IOWait
	bIdle := statB.Idle + statB.IOWait

	aNonIdle := statA.User + statA.Nice + statA.System + statA.IRQ + statA.SoftIRQ + statA.Steal
	bNonIdle := statB.User + statB.Nice + statB.System + statB.IRQ + statB.SoftIRQ + statB.Steal

	aTotal := aIdle + aNonIdle
	bTotal := bIdle + bNonIdle

	totalDiff := bTotal - aTotal
	idleDiff := bIdle - aIdle

	return uint((float64(totalDiff-idleDiff) / float64(totalDiff)) * 100)
}
