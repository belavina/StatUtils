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
func sysstatsHandler(w http.ResponseWriter, r *http.Request) {
	jsonData := perfstats.PlatformSysStats()
	w.Write(jsonData)
}

// starting point
func main() {

	// -- define command line args
	httpPortPtr := flag.Int("port", 9159, "HTTP server port")
	flag.Parse()

	perfstats.SayHi()

	// start up http server
	fmt.Printf("Listening on port %d\n", *httpPortPtr)

	http.HandleFunc("/sysstats", sysstatsHandler)
	http.ListenAndServe(fmt.Sprintf(":%d", *httpPortPtr), nil)
}
