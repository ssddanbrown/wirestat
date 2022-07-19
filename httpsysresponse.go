package main

import (
	"fmt"
	"os"
)

func main() {

	fileSystems, err := GetFileSystemMap()
	if err != nil {
		panic(err)
	}

	cpus, err := GetCpuMap()
	if err != nil {
		panic(err)
	}

	memory, err := GetMemory()
	if err != nil {
		panic(err)
	}

	uptime, err := GetUptime()
	if err != nil {
		panic(err)
	}

	dd(fileSystems, cpus, memory, uptime)
}

func dd(data ...interface{}) {
	for _, item := range data {
		fmt.Printf("%+v\n", item)
	}
	os.Exit(1)
}
