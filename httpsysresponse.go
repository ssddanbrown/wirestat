package main

import (
	"fmt"
	"os"
)

func main() {
	rules, err := parseRuleFile("/home/dan/GolandProjects/httpsysresponse/rules.txt")
	if err != nil {
		panic(err)
	}

	responseBuilder := NewResponseBuilder(rules)

	go StartPollingSystem()
	startServer(responseBuilder)
}

func dd(data ...interface{}) {
	for _, item := range data {
		fmt.Printf("%+v\n", item)
	}
	os.Exit(1)
}
