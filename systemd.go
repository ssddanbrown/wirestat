package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func GenerateSystemdConfig(port uint, rulesPath string) string {

	execPath, err := os.Executable()
	if err != nil {
		log.Panicf("Failed getting the current executable path with error: %s", err.Error())
	}

	config := fmt.Sprintf(`
[Unit]
Description=WireStat
After=network.target

[Service]
Type=simple
Restart=always
RestartSec=1s
ExecStart=%s -port %d -rules "%s"

[Install]
WantedBy=multi-user.target
	`, execPath, port, rulesPath)

	return strings.TrimSpace(config)
}
