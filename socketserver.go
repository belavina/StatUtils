package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strings" // only needed below for sample processing
)

// Performance Monitor
// relog 'C:\Path' -o text.csv -f csv

func createCollector(collectorSet string, logfilePath string) {
	fmt.Printf("[task] Creating collector '%s'\n", collectorSet)

	logmanCmd := exec.Command(
		"logman.exe",
		"create",
		"counter",
		collectorSet,
		"-c",
		"\"\\LogicalDisk(*)\\% Free Space\"",
		"-f",
		"csv",
		"-ow", // overwrite
		"-o",
		logfilePath)

	out, err := logmanCmd.Output()
	if err != nil {
		fmt.Printf("%s", out)
		fmt.Println(err)
	}
}

func startCollector(collectorSet string) {

	fmt.Printf("[task] Starting collector '%s'\n", collectorSet)

	logmanCmd := exec.Command(
		"logman.exe",
		"start",
		collectorSet)

	out, err := logmanCmd.Output()
	if err != nil {
		fmt.Printf("%s", out)
		fmt.Println(err)
	}
}

func launchServer() {
	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8829")

	// accept connection on port
	conn, _ := ln.Accept()

	// run loop forever (or until ctrl-c)
	for {
		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// output message received
		fmt.Print("Message Received:", string(message))
		// sample process for string received
		newmessage := strings.ToUpper(message)
		// send new string back to client
		conn.Write([]byte(newmessage + "\n"))
	}
}

func getCollectorStatus(collectorSet string) map[string]string {

	logmanCmd := exec.Command("logman.exe", "query", collectorSet)
	out, err := logmanCmd.Output()
	if err != nil {
		fmt.Printf("%s", out)
		fmt.Println(err)
		return nil
	}

	var collectorStatusMap map[string]string

	collectorStatusMap = make(map[string]string)

	outStr := string(out[:])
	lines := strings.Split(outStr, "\n")
	for _, line := range lines {
		keyValuePairs := strings.Split(line, ":")
		if len(keyValuePairs) < 2 {
			continue
		}

		value := strings.TrimSpace(strings.Join(keyValuePairs[1:], ":"))
		collectorStatusMap[keyValuePairs[0]] = value

	}

	return collectorStatusMap
}

func getCurrentLogFile(collectorSet string) string {
	logFile := getCollectorStatus(collectorSet)["Output Location"]

	return logFile
}

type SysStat struct {
	date  string
	key   string
	value string
}

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

	for _, each := range csvData {
		statEntry.date = each[0]
		statEntry.key = each[1]
		statEntry.value = each[2]

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
	iniCollectorPtr := flag.Bool("ini-collector", false, "Initialize logman collector")
	startCollectorPtr := flag.Bool("start-collector", false, "Start perf-agent collector")
	runAgentPtr := flag.Bool("start-agent", false, "Start a server for the agent")
	logfilePath := flag.String("log-location", "c:\\perflogs\\anvil_agent_stats.csv", "Anvil agent log location")
	collectorSet := flag.String("collector-name", "anvil_agent_win", "Anvil data collector name")

	flag.Parse()

	// cli callbacks
	if *iniCollectorPtr {
		createCollector(*collectorSet, *logfilePath)
	}

	if *startCollectorPtr {
		startCollector(*collectorSet)
	}

	if *runAgentPtr {

		// out := getCurrentLogFile(*collectorSet)
		// fmt.Println(out)
		http.HandleFunc("/sysstats", sysstatsHandler)
		http.ListenAndServe(":9159", nil)
		// launchServer()
	}
}
