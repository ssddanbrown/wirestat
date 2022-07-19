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

	dd(fileSystems)
}

func dd(data interface{}) {
	fmt.Printf("%+v\n", data)
	os.Exit(1)
}
