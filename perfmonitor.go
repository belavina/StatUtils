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
	Status  string `json:"status"`
	Data    []byte `json:"data"`
	Message string `json:"message"`
}

type appHandler func(http.ResponseWriter, *http.Request) error

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		http.Error(w, err.Error(), 500)
	}
}

// Processes http request for latest system performance statistics
// func sysStatsHandler(w http.ResponseWriter, r *http.Request) {

// 	// json.Marshal(stats)
// 	w.Write(json.Marshal(perfstats.PlatformSysStats()))
// }

// Get host details such as platform & hostname
func computerInfoHandler(w http.ResponseWriter, r *http.Request) error {
	data, err := perfstats.GetPlatformInfo()

	if err != nil {
		return err
	}

	w.Write(json.Marshal(data))
	return nil
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
	// http.HandleFunc("/sysstats", appHandler(sysStatsHandler))
	http.HandleFunc("/platform", appHandler(computerInfoHandler))

	http.ListenAndServe(fmt.Sprintf(":%d", *httpPortPtr), nil)
}
