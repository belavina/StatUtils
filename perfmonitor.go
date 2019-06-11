// Created by Olga Belavina
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"net/http"

	"./perfstats"
)

// Processes http request for latest system performance statistics
func sysStatsHandler(w http.ResponseWriter, r *http.Request) {
	jsonData := perfstats.PlatformSysStats()
	w.Write(jsonData)
}

// Get host details such as platform & hostname
func computerInfoHandler(w http.ResponseWriter, r *http.Request) {
	jsonData := perfstats.GetPlatformInfo()
	w.Write(jsonData)
}

func main() {

	// -- define command line args
	httpPortPtr := flag.Int("port", 9159, "HTTP server port")
	flag.Parse()

	// start up http server
	fmt.Printf("Listening on port %d\n", *httpPortPtr)

	// http routes:
	http.HandleFunc("/sysstats", sysStatsHandler)
	http.HandleFunc("/info", computerInfoHandler)

	http.ListenAndServe(fmt.Sprintf(":%d", *httpPortPtr), nil)
}
