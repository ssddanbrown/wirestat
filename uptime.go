package main

import (
	"fmt"
	linuxproc "github.com/c9s/goprocinfo/linux"
	"math"
)

type Uptime struct {
	seconds int
	minutes int
	hours   int
	days    int
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
		seconds: seconds,
		minutes: minutes,
		hours:   hours,
		days:    days,
	}

	return uptime, nil
}
