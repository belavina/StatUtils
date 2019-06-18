package perfstats

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"os/exec"
)

// Convert sysStat in csv format to json
func sysStatCSVToSysStat(cmdOut []byte) ([]SysStat, error) {

	var statEntry SysStat
	var stats []SysStat

	reader := csv.NewReader(bytes.NewReader(cmdOut))
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()

	if err != nil {
		return stats, fmt.Errorf("Error while parsing .csv script output: %s", err)
	}

	if csvData == nil {
		return stats, errors.New("CSV data is empty")
	}

	for _, each := range csvData[1:] {
		statEntry.Date = each[0]
		statEntry.Key = each[1]
		statEntry.Value = each[2]

		stats = append(stats, statEntry)
	}

	return stats, nil
}

// Query performance stats on windows platform
func queryWindowsSysStats() ([]byte, error) {

	cmdResult := exec.Command("powershell.exe", "-executionpolicy", "bypass", "-file", "./SysStats.ps1")

	out, err := cmdResult.Output()
	if err != nil {
		return nil, err
	}

	return out, nil
}

func PlatformSysStats() (interface{}, error) {

	csvOutput, err := queryWindowsSysStats()
	if err != nil {
		return nil, fmt.Errorf("Cannot get windows system performance details: %s", err)
	}

	winStats, err := sysStatCSVToSysStat(csvOutput)
	if err != nil {
		return nil, fmt.Errorf("Cannot parse SysStats ps .csv: output %s", err)
	}

	return winStats, nil
}
