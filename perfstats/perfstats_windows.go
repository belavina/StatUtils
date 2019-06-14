package perfstats

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os/exec"
)

// Convert sysStat in csv format to json
func sysStatCSVToSysStat(cmdOut []byte) SysStat {

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

	cmdResult := exec.Command("powershell.exe", "-executionpolicy", "bypass", "-file", "./SysStats.ps1")

	out, err := cmdResult.Output()
	if err != nil {
		fmt.Println(err)
	}

	return out
}

func PlatformSysStats() (interface{}, error) {
	return sysStatCSVToSysStat(queryWindowsSysStats()), nil
}
