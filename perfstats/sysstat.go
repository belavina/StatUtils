package perfstats

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
)

// SysStat Performance statistics
type SysStat struct {
	Date  string // date when stat item was grabbed
	Key   string // type of sys stats (memory, cpu etc.)
	Value string // current value of sys stat (cpu usage %, free space)
}

// GetPlatformInfo show platform details
func GetPlatformInfo() []byte {

	hostname, _ := os.Hostname()
	jsonData, err := json.Marshal(map[string]string{
		"platform": runtime.GOOS,
		"machine":  hostname,
	})

	if err != nil {
		fmt.Println("Error!")
		fmt.Println(err)
	}

	return jsonData
}
