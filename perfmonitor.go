// Created by Olga Belavina
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/Seneca-CDOT/StatUtils/perfstats"
)

// AppVersion - current app version
const AppVersion = "0.0.1"

type response struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type getResponseData func() (interface{}, error)

func processRequest(w http.ResponseWriter, getData getResponseData) {
	w.Header().Set("Content-Type", "application/json")

	data, err := getData()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response{
			Status:  "error",
			Data:    nil,
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response{
		Status:  "success",
		Data:    data,
		Message: "",
	})
}

// Processes http request for latest system performance statistics
func sysStatsHandler(w http.ResponseWriter, r *http.Request) {
	processRequest(w, perfstats.PlatformSysStats)
}

// Get host details such as platform & hostname
func computerInfoHandler(w http.ResponseWriter, r *http.Request) {
	processRequest(w, perfstats.GetPlatformInfo)
}

func main() {

	// -- define command line args
	httpPortPtr := flag.Int("port", 9159, "HTTP server port")
	version := flag.Bool("version", false, "print current perfmonitor version and exit")

	flag.Parse()

	if *version {
		fmt.Println(AppVersion)
		os.Exit(0)
	}

	// start up http server
	fmt.Printf("Listening on port %d\n", *httpPortPtr)

	// http routes:
	http.HandleFunc("/sysstats", sysStatsHandler)
	http.HandleFunc("/platform", computerInfoHandler)

	http.ListenAndServe(fmt.Sprintf(":%d", *httpPortPtr), nil)
}
