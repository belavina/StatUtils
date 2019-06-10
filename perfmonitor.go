// Created by Olga Belavina
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

// AppVersion - current app version
const AppVersion = "1.0.0"

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

// Processes http request for latest system performance statistics
func sysstatsHandler(w http.ResponseWriter, r *http.Request) {
	jsonData := sysStatCSVToJSON(queryWindowsSysStats())
	w.Write(jsonData)
}

// starting point
func main() {

	// -- define command line args
	httpPortPtr := flag.Int("port", 9159, "HTTP server port")
	version := flag.Bool("version", false, "prints current perfmonitor versio and exit")

	flag.Parse()

	if *version {
		fmt.Println(AppVersion)
		os.Exit(0)
	}

	// start up http server
	fmt.Printf("Listening on port %d\n", *httpPortPtr)

	http.HandleFunc("/sysstats", sysstatsHandler)
	http.ListenAndServe(fmt.Sprintf(":%d", *httpPortPtr), nil)
}
