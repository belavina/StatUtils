package perfstats

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"os/exec"
)

// Convert sysStat in csv format to map
func csvToMap(cmdOut []byte) ([]map[string]string, error) {

	// var statEntry SysStat
	var stats []map[string]string

	reader := csv.NewReader(bytes.NewReader(cmdOut))
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()

	if err != nil {
		return stats, fmt.Errorf("Error while parsing .csv script output: %s", err)
	}

	if csvData == nil {
		return stats, errors.New("CSV data is empty")
	}

	headers := csvData[0]

	for _, each := range csvData[1:] {

		statEntry := make(map[string]string)
		for i := range headers {
			statEntry[headers[i]] = each[i]
		}
		stats = append(stats, statEntry)
	}

	return stats, nil
}

func getPerfCounter(counterName string) (StatEntry, error) {
	var statEntry StatEntry
	statEntry.Date = getDateFormatted()

	getCounterFmt := "& {Get-Counter -Counter \"%s\" | Select-Object -ExpandProperty CounterSamples | convertto-csv -NoTypeInformation}"
	getCounter := fmt.Sprintf(getCounterFmt, counterName)
	cmdResult := exec.Command("powershell.exe", "-Command", getCounter)

	out, err := cmdResult.Output()

	if err != nil {
		return statEntry, err
	}

	// Convert to csv
	statEntry.Stats, err = csvToMap(out)
	if err != nil {
		return statEntry, err
	}

	return statEntry, nil
}

func getCPUStats() (StatEntry, error) {
	return getPerfCounter("\\Processor Information(*)\\% Processor Time")
}

func getDiskStats() (StatEntry, error) {
	return getPerfCounter("\\LogicalDisk(*)\\% Free Space")
}

func getMemoryStats() (StatEntry, error) {
	return getPerfCounter("\\Memory\\Available Bytes")
}
