package main

import (
	"bufio"
	"bytes"
	"os/exec"
	"strconv"
	"strings"
)

type FileSystemMap map[string]*FileSystem

type FileSystem struct {
	name        string
	capacity    uint64
	used        uint64
	usedPercent uint
	available   uint64
}

func GetFileSystemMap() (FileSystemMap, error) {
	dfOutput, err := getDfOutput()
	if err != nil {
		return nil, err
	}

	fileSystems, err := parseDfOutputToFileSystem(dfOutput)
	if err != nil {
		return nil, err
	}

	fsMap := make(FileSystemMap)
	for _, fileSystem := range fileSystems {
		fsMap[fileSystem.name] = fileSystem
	}

	return fsMap, nil
}

func getDfOutput() ([]byte, error) {
	return exec.Command("df", "-P", "-B", "MB").Output()
}

func parseDfOutputToFileSystem(dfOutput []byte) ([]*FileSystem, error) {
	fileSystems := []*FileSystem{}
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

func parseDfLineToFileSystem(line string) (*FileSystem, error) {
	fileSystem := FileSystem{}
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
