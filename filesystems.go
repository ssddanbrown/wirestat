package main

import (
	"bufio"
	"bytes"
	"os/exec"
	"strconv"
	"strings"
)

type fileSystem struct {
	name        string
	capacity    uint64
	used        uint64
	usedPercent uint
	available   uint64
}

func GetFileSystemMap() (map[string]uint64, error) {
	dfOutput, err := getDfOutput()
	if err != nil {
		return nil, err
	}

	fileSystems, err := parseDfOutputToFileSystem(dfOutput)
	if err != nil {
		return nil, err
	}

	fsMap := make(map[string]uint64)
	for _, fileSystem := range fileSystems {
		fsMap[fileSystem.name+".capacity"] = fileSystem.capacity
		fsMap[fileSystem.name+".used"] = fileSystem.used
		fsMap[fileSystem.name+".used_percent"] = uint64(fileSystem.usedPercent)
		fsMap[fileSystem.name+".available"] = fileSystem.available
	}

	return fsMap, nil
}

func getDfOutput() ([]byte, error) {
	return exec.Command("df", "-P", "-B", "MB").Output()
}

func parseDfOutputToFileSystem(dfOutput []byte) ([]*fileSystem, error) {
	fileSystems := []*fileSystem{}
	scanner := bufio.NewScanner(bytes.NewReader(dfOutput))
	lineCount := 0

	for scanner.Scan() {
		lineCount++
		if lineCount == 1 {
			continue
		}

		line := scanner.Text()
		fileSystem, err := parseDfLineToFileSystem(line)
		if err != nil {
			return nil, err
		}

		fileSystems = append(fileSystems, fileSystem)
	}

	return fileSystems, nil
}

func parseDfLineToFileSystem(line string) (*fileSystem, error) {
	fileSystem := fileSystem{}
	parts := strings.Split(line, " ")
	index := 0

	for _, part := range parts {
		if part == "" {
			continue
		}

		if index == 0 {
			fileSystem.name = part
		}
		if index == 2 {
			i, err := strconv.ParseUint(strings.TrimRight(part, "MB"), 10, 64)
			if err != nil {
				return nil, err
			}
			fileSystem.used = i
		}
		if index == 3 {
			i, err := strconv.ParseUint(strings.TrimRight(part, "MB"), 10, 64)
			if err != nil {
				return nil, err
			}
			fileSystem.available = i
		}
		if index == 4 {
			percent := strings.TrimRight(part, "%")
			i, err := strconv.ParseUint(percent, 10, 32)
			if err != nil {
				return nil, err
			}
			fileSystem.usedPercent = uint(i)
		}

		index++
	}

	fileSystem.capacity = fileSystem.used + fileSystem.available
	return &fileSystem, nil
}
