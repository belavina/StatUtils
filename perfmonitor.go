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
const AppVersion = "0.1.0"

type response struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type appHandler func(http.ResponseWriter, *http.Request) (interface{}, error)

// implement the http.Handler interface's ServeHTTP to add error handling
// and uniform json response format
func (getData appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := getData(w, r)

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
func sysStatsHandler(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return perfstats.PlatformSysStats()
}

// Get host details such as platform & hostname
func computerInfoHandler(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return perfstats.GetPlatformInfo()
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
	http.Handle("/sysstats", appHandler(sysStatsHandler))
	http.Handle("/platform", appHandler(computerInfoHandler))

	http.ListenAndServe(fmt.Sprintf(":%d", *httpPortPtr), nil)
}
