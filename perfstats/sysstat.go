package perfstats

import (
	"fmt"
	"os"
	"runtime"
	"strings"
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

// Web API date format (utc timestamp)
func getDateFormatted() string {
	return time.Now().UTC().Format("20060102150405")
}

// Lowers only the first character in a string
func lowerFirst(str string) string {
	return strings.ToLower(string(str[0])) + str[1:]
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

// PlatformSysStats Query performance stats
// (calls either _linux.go or _windows perfstats implementations)
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
