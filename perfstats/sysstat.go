package perfstats

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

// AppVersion - current app version
const AppVersion = "0.2.0"

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

type statGetter func() (StatEntry, error)

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

	hostname, err := os.Hostname()
	machineDetails := map[string]string{
		"platform":        runtime.GOOS,
		"machine":         hostname,
		"softwareVersion": AppVersion,
	}
	return machineDetails, err
}

// PlatformSysStats Query performance stats
// (calls either _linux.go or _windows perfstats implementations)
func PlatformSysStats() (interface{}, error) {

	statEntries := make(map[string]StatEntry)
	perfStatGetters := map[string]statGetter{
		"memory": getMemoryStats,
		"cpu":    getCPUStats,
		"disk":   getDiskStats,
	}

	for statName, getStatsFunc := range perfStatGetters {
		stats, err := getStatsFunc()
		if err != nil {
			return nil, fmt.Errorf("Cannot get %s details: %s", statName, err)
		}
		stats.Date = getDateFormatted()
		statEntries[statName] = stats
	}

	return SysStat{
		Memory: statEntries["memory"],
		CPU:    statEntries["cpu"],
		Disk:   statEntries["disk"],
	}, nil
}
