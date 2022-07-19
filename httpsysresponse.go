package main

import (
	"fmt"
	"os"
)

func main() {

	disks, err := GetDiskMap()
	if err != nil {
		panic(err)
	}

	dd(disks)
}

func dd(data interface{}) {
	fmt.Printf("%+v\n", data)
	os.Exit(1)
}
