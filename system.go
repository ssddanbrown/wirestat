package main

import (
	"fmt"
	"time"
)

type System struct {
	Filesystem       map[string]uint64
	Cpu              map[string]uint64
	Memory           map[string]uint64
	Uptime           map[string]uint64
	MetricsUpdatedAt time.Time
}

func (s *System) mergeMaps() map[string]uint64 {
	merged := make(map[string]uint64)
	maps := []map[string]uint64{s.Filesystem, s.Cpu, s.Memory, s.Uptime}
	prefixes := []string{"filesystem", "cpu", "memory", "uptime"}

	for mapIndex, sMap := range maps {
		for k, v := range sMap {
			fullKey := prefixes[mapIndex] + "." + k
			merged[fullKey] = v
		}
	}

	return merged
}

var latestSystem *System

func GetLatestSystem() *System {
	return latestSystem
}

func StartPollingSystem() {
	for {
		system, err := getSystem()

		if err != nil {
			panic(fmt.Sprintf("Failed loading system data with error: %s", err))
		}

		latestSystem = system
		time.Sleep(time.Second * 5)
	}
}

func getSystem() (*System, error) {

	fileSystems, err := GetFileSystemMap()
	if err != nil {
		return nil, err
	}

	cpus, err := GetCpuMap()
	if err != nil {
		return nil, err
	}

	memory, err := GetMemory()
	if err != nil {
		return nil, err
	}

	uptime, err := GetUptime()
	if err != nil {
		return nil, err
	}

	system := &System{
		Filesystem:       fileSystems,
		Cpu:              cpus,
		Memory:           memory,
		Uptime:           uptime,
		MetricsUpdatedAt: time.Now(),
	}

	return system, nil
}
