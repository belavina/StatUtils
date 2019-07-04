package perfstats

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"os/exec"
)

// Convert csv to array of maps
func parseCSVOutput(cmdOut []byte) ([]map[string]string, error) {

	var parsedCsv []map[string]string

	reader := csv.NewReader(bytes.NewReader(cmdOut))
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()

	if err != nil {
		return parsedCsv, fmt.Errorf("Error while parsing .csv script output: %s", err)
	}

	if csvData == nil {
		return parsedCsv, errors.New("CSV data is empty")
	}

	headers := csvData[0]
	// lowercase first letter and match value format with other platforms
	for idx, header := range headers {
		if header == "CookedValue" {
			header = "value"
		}
		headers[idx] = lowerFirst(header)
	}

	// parse rows of .csv data
	for _, each := range csvData[1:] {

		csvEntry := make(map[string]string)

		for i := range headers {
			csvEntry[headers[i]] = each[i]
		}
		parsedCsv = append(parsedCsv, csvEntry)
	}

	return parsedCsv, nil
}

// Query windows counters with powershell command
func getPerfCounter(counterName string) (StatEntry, error) {
	var statEntry StatEntry

	// Run powershell command returning a performance counter
	getCounterFmt := "& {Get-Counter -Counter \"%s\" | Select-Object -ExpandProperty CounterSamples | convertto-csv -NoTypeInformation}"
	getCounter := fmt.Sprintf(getCounterFmt, counterName)
	cmdResult := exec.Command("powershell.exe", "-executionpolicy", "bypass", "-Command", getCounter)

	out, err := cmdResult.Output()

	if err != nil {
		return statEntry, err
	}

	// Convert to csv
	statEntry.Stats, err = parseCSVOutput(out)
	if err != nil {
		return statEntry, err
	}

	return statEntry, nil
}

func getCPUStats() (StatEntry, error) {
	return getPerfCounter("\\Processor Information(*)\\% Processor Time")
}

func getDiskStats() (StatEntry, error) {

	var statEntry StatEntry

	// Query logical disk restricting drives to "local disk" type (#3)
	getWmiCommand := "& {Get-WmiObject -Class Win32_logicaldisk -Filter \"DriveType = '3'\" "
	// Get total, free and calculate used (all in bytes)
	getWmiCommand += "| Select-Object -Property DeviceID, Size, FreeSpace, @{L=\"Used\";E={\"{0}\" -f ($_.Size-$_.FreeSpace)} } "
	getWmiCommand += "| convertto-csv -NoTypeInformation}"
	cmdResult := exec.Command("powershell.exe", "-executionpolicy", "bypass", "-Command", getWmiCommand)

	out, err := cmdResult.Output()

	if err != nil {
		return statEntry, err
	}

	// Convert to csv
	statEntry.Stats, err = parseCSVOutput(out)
	if err != nil {
		return statEntry, err
	}

	return statEntry, nil

}

func getMemoryStats() (StatEntry, error) {
	return getPerfCounter("\\Memory\\Available Bytes")
}
