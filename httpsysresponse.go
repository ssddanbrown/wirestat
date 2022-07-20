package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	portOpt := flag.Uint("port", 8930, "Port to run the server on")
	rulesPathOpt := flag.String("rules", "/etc/httpsysresponse/rules.txt", "Path to the file containing rules")

	flag.Parse()

	_, err := os.Stat(*rulesPathOpt)
	if err != nil {
		fmt.Println(fmt.Sprintf("Startup failed, error when checking rules file at %s", *rulesPathOpt))
		os.Exit(1)
	}

	rules, err := parseRuleFile(*rulesPathOpt)
	if err != nil {
		fmt.Println(fmt.Sprintf("Startup failed, error when parsing rules file: %s", err.Error()))
		os.Exit(1)
	}

	responseBuilder := NewResponseBuilder(rules)

	go StartPollingSystem()

	fmt.Println(fmt.Sprintf("Starting server, listening on: http://0.0.0.0:%d", *portOpt))
	startServer(responseBuilder, *portOpt)
}

func dd(data ...interface{}) {
	for _, item := range data {
		fmt.Printf("%+v\n", item)
	}
	os.Exit(1)
}
