package main

type System struct {
	Alerts     []string      `json:"alerts"`
	Filesystem FileSystemMap `json:"filesystem"`
	Cpu        CpuMap        `json:"cpu"`
	Memory     *Memory       `json:"memory"`
	Uptime     *Uptime       `json:"uptime"`
}

func getSystem() (*System, error) {

	fileSystems, err := GetFileSystemMap()
	if err != nil {
		return nil, err
	}

	cpus, err := GetCpuMap()
	if err != nil {
		return nil, err
	}

	memory, err := GetMemory()
	if err != nil {
		return nil, err
	}

	uptime, err := GetUptime()
	if err != nil {
		return nil, err
	}

	system := &System{
		Alerts:     []string{},
		Filesystem: fileSystems,
		Cpu:        cpus,
		Memory:     memory,
		Uptime:     uptime,
	}

	return system, nil
}
