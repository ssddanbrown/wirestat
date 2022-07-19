package main

import (
	"fmt"
	"os"
)

func main() {

	//fileSystems, err := GetFileSystemMap()
	//if err != nil {
	//	panic(err)
	//}
	//
	//dd(fileSystems)

	//cpus, err := GetCpuMap()
	//if err != nil {
	//	panic(err)
	//}
	//
	//dd(cpus)

	_, err := GetMemory()
	if err != nil {
		panic(err)
	}

	//dd(memory)

	uptime, err := GetUptime()
	if err != nil {
		panic(err)
	}

	dd(uptime)
}

func dd(data interface{}) {
	fmt.Printf("%+v\n", data)
	os.Exit(1)
}
