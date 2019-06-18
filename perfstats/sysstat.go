package perfstats

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

// StatEntry specific performance data sweep
type StatEntry struct {
	Stats interface{} `json:"stats"`
	Date  string      `json:"date"`
}

// SysStat Performance statistics
type SysStat struct {
	CPU    StatEntry `json:"cpu"`
	Disk   StatEntry `json:"disk"`
	Memory StatEntry `json:"memory"`
}

func getDateFormatted() string {
	dt := time.Now()
	return dt.Format("1/2/2006 3:04:05 PM")
}

// GetPlatformInfo show platform details
func GetPlatformInfo() (interface{}, error) {

	hostname, _ := os.Hostname()
	machineDetails := map[string]string{
		"platform": runtime.GOOS,
		"machine":  hostname,
	}
	return machineDetails, nil
}

// PlatformSysStats Query performance stats on linux platform
func PlatformSysStats() (interface{}, error) {

	memInfo, err := getMemoryStats()
	if err != nil {
		return nil, fmt.Errorf("Cannot get memory details: %s", err)
	}
	cpuInfo, err := getCPUStats()
	if err != nil {
		return nil, fmt.Errorf("Cannot get CPU details: %s", err)
	}
	diskInfo, err := getDiskStats()
	if err != nil {
		return nil, fmt.Errorf("Cannot get disk details: %s", err)
	}

	return SysStat{
		Memory: memInfo,
		CPU:    cpuInfo,
		Disk:   diskInfo,
	}, nil
}
