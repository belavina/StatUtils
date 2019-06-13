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
func GetPlatformInfo() (map[string]string, error) {

	hostname, _ := os.Hostname()
	return map[string]string{
		"platform": runtime.GOOS,
		"machine":  hostname,
	}, nil
}
