package perfstats

import (
	"os"
	"runtime"
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

// GetPlatformInfo show platform details
func GetPlatformInfo() (interface{}, error) {

	hostname, _ := os.Hostname()
	machineDetails := map[string]string{
		"platform": runtime.GOOS,
		"machine":  hostname,
	}
	return machineDetails, nil
}
