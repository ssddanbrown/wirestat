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

	memory, err := GetMemory()
	if err != nil {
		panic(err)
	}

	dd(memory)
}

func dd(data interface{}) {
	fmt.Printf("%+v\n", data)
	os.Exit(1)
}
