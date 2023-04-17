package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func GenerateSystemdConfig(port uint, rulesPath string, accessKey string, ruleDelimeter string) string {

	execPath, err := os.Executable()
	if err != nil {
		log.Panicf("Failed getting the current executable path with error: %s", err.Error())
	}

	accessKeyOpt := ""
	if accessKey != "" {
		accessKeyOpt = fmt.Sprintf(`-accesskey "%s"`, accessKey)
	}

	ruleDelimeterOpt := ""
	if ruleDelimeter != ":" {
		ruleDelimeterOpt = fmt.Sprintf(`-ruledelimeter "%s"`, ruleDelimeter)
	}

	config := fmt.Sprintf(`
[Unit]
Description=wirestat
After=network.target

[Service]
Type=simple
Restart=always
RestartSec=1s
ExecStart=%s -port %d -rules "%s" %s %s

[Install]
WantedBy=multi-user.target
	`, execPath, port, rulesPath, accessKeyOpt, ruleDelimeterOpt)

	return strings.TrimSpace(config)
}
