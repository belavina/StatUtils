package perfstats

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os/exec"
)

func SayHi() {
	fmt.Println("hi windows")
}

// SysStat Performance statistics
type SysStat struct {
	Date  string // date when stat item was grabbed
	Key   string // type of sys stats (memory, cpu etc.)
	Value string // current value of sys stat (cpu usage %, free space)
}

// Convert sysStat in csv format to json
func sysStatCSVToJSON(cmdOut []byte) []byte {

	reader := csv.NewReader(bytes.NewReader(cmdOut))
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()

	if csvData == nil {
		fmt.Println("CSV data is empty")
		return nil
	}

	if err != nil {
		fmt.Println(err)
		return nil
	}

	var statEntry SysStat
	var stats []SysStat

	for _, each := range csvData[1:] {
		statEntry.Date = each[0]
		statEntry.Key = each[1]
		statEntry.Value = each[2]

		stats = append(stats, statEntry)
	}

	jsonData, err := json.Marshal(stats)

	if err != nil {
		fmt.Println(err)
	}

	return jsonData
}

// Query performance stats on windows platform
func queryWindowsSysStats() []byte {

	// sysStats := exec.Command("/bin/bash", "./fake.sh")
	sysStats := exec.Command("powershell.exe", "./SysStats.ps1")

	out, err := sysStats.Output()
	if err != nil {
		fmt.Println(err)
	}

	return out
}

func PlatformSysStats() []byte {
	return sysStatCSVToJSON(queryWindowsSysStats())
}
