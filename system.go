package main

import (
	"fmt"
	"time"
)

type System struct {
	Alerts         []string      `json:"alerts"`
	Filesystem     FileSystemMap `json:"filesystem"`
	Cpu            CpuMap        `json:"cpu"`
	Memory         *Memory       `json:"memory"`
	Uptime         *Uptime       `json:"uptime"`
	StatsUpdatedAt time.Time     `json:"stats_updated_at"`
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
		Alerts:         []string{},
		Filesystem:     fileSystems,
		Cpu:            cpus,
		Memory:         memory,
		Uptime:         uptime,
		StatsUpdatedAt: time.Now(),
	}

	return system, nil
}
