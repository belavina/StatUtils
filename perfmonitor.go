// Created by Olga Belavina
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
)

// SysStat Performance statistics
type SysStat struct {
	Date  string // date when stat item was grabbed
	Key   string // type of sys stats (memory, cpu etc.)
	Value string // current value of sys stat (cpu usage %, free space)
}

// Processes http request for latest system performance statistics
func sysstatsHandler(w http.ResponseWriter, r *http.Request) {

	csvFile, err := os.Open("./test.csv")
	if err != nil {
		fmt.Println(err)
	}

	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		return
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

	w.Write(jsonData)
}

// starting point
func main() {

	// -- define command line args
	httpPortPtr := flag.Int("port", 9159, "HTTP server port")
	flag.Parse()

	// start up http server
	fmt.Printf("Listening on port %d\n", *httpPortPtr)

	http.HandleFunc("/sysstats", sysstatsHandler)
	http.ListenAndServe(fmt.Sprintf(":%d", *httpPortPtr), nil)
}
