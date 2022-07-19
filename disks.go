package main

import (
	"bufio"
	"bytes"
	"os/exec"
	"strconv"
	"strings"
)

type DiskPartitionMap map[string]*DiskPartition

type DiskPartition struct {
	fileSystem  string
	capacity    uint64
	used        uint64
	usedPercent uint
	available   uint64
}

func GetDiskMap() (DiskPartitionMap, error) {
	dfOutput, err := getDfOutput()
	if err != nil {
		return nil, err
	}

	partitions, err := parseDfOutputToPartition(dfOutput)
	if err != nil {
		return nil, err
	}

	diskMap := make(DiskPartitionMap)
	for _, partition := range partitions {
		diskMap[partition.fileSystem] = partition
	}

	return diskMap, nil
}

func getDfOutput() ([]byte, error) {
	return exec.Command("df", "-P").Output()
}

func parseDfOutputToPartition(dfOutput []byte) ([]*DiskPartition, error) {
	partitions := []*DiskPartition{}
	scanner := bufio.NewScanner(bytes.NewReader(dfOutput))
	lineCount := 0

	for scanner.Scan() {
		lineCount++
		if lineCount == 1 {
			continue
		}

		line := scanner.Text()
		partition, err := parseDfLineToPartition(line)
		if err != nil {
			return nil, err
		}

		partitions = append(partitions, partition)
	}

	return partitions, nil
}

func parseDfLineToPartition(line string) (*DiskPartition, error) {
	partition := DiskPartition{}
	parts := strings.Split(line, " ")
	index := 0

	for _, part := range parts {
		if part == "" {
			continue
		}

		if index == 0 {
			partition.fileSystem = part
		}
		if index == 2 {
			i, err := strconv.ParseUint(part, 10, 64)
			if err != nil {
				return nil, err
			}
			partition.used = i
		}
		if index == 3 {
			i, err := strconv.ParseUint(part, 10, 64)
			if err != nil {
				return nil, err
			}
			partition.available = i
		}
		if index == 4 {
			percent := strings.TrimRight(part, "%")
			i, err := strconv.ParseUint(percent, 10, 32)
			if err != nil {
				return nil, err
			}
			partition.usedPercent = uint(i)
		}

		index++
	}

	partition.capacity = partition.used + partition.available
	return &partition, nil
}
