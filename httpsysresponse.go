package main

import (
	"fmt"
	"os"
)

func main() {
	startServer()
}

func dd(data ...interface{}) {
	for _, item := range data {
		fmt.Printf("%+v\n", item)
	}
	os.Exit(1)
}
