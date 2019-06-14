package perfstats

import (
	"os"
	"runtime"
)

// SysStat Performance statistics
type SysStat struct {
	Date  string `json:"date"`
	Key   string `json:"key"`
	Value string `json:"value"`
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
