// Created by Olga Belavina
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"./perfstats"
)

// AppVersion - current app version
const AppVersion = "1.0.0"

// SysStat Performance statistics
type SysStat struct {
	Date  string // date when stat item was grabbed
	Key   string // type of sys stats (memory, cpu etc.)
	Value string // current value of sys stat (cpu usage %, free space)
}

// Processes http request for latest system performance statistics
func sysStatsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(perfstats.PlatformSysStats())
}

// Get host details such as platform & hostname
func computerInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(perfstats.GetPlatformInfo())
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
