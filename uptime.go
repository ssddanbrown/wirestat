package main

import (
	"fmt"
	linuxproc "github.com/c9s/goprocinfo/linux"
	"math"
)

type Uptime struct {
	Seconds int `json:"seconds"`
	Minutes int `json:"minutes"`
	Hours   int `json:"hours"`
	Days    int `json:"days"`
}

func GetUptime() (*Uptime, error) {
	utInfo, err := linuxproc.ReadUptime("/proc/uptime")
	if err != nil {
		return nil, fmt.Errorf("failed to stat uptime data, received error: %s", err.Error())
	}

	seconds := int(utInfo.Total)
	days := int(math.Floor(float64(seconds) / 86400))
	seconds = seconds - (days * 86400)

	hours := int(math.Floor(float64(seconds) / 3600))
	seconds = seconds - (hours * 3600)

	minutes := int(math.Floor(float64(seconds) / 60))
	seconds = seconds - (minutes * 60)

	uptime := &Uptime{
		Seconds: seconds,
		Minutes: minutes,
		Hours:   hours,
		Days:    days,
	}

	return uptime, nil
}
