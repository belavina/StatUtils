package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os/exec"
	"strings" // only needed below for sample processing
)

// Performance Monitor
// relog 'C:\Path' -o text.csv -f csv

func createCollector(collectorSet string, logfilePath string) {
	// counters = ("\\LogicalDisk(*)\\% Free Space")

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

// starting point
func main() {

	// -- define command line args
	iniCollectorPtr := flag.Bool("ini-collector", false, "Initialize logman collector")
	startCollectorPtr := flag.Bool("start-collector", false, "Start perf-agent collector")
	runAgentPtr := flag.Bool("start-agent", false, "Start a server for the agent")

	flag.Parse()

	logfilePath := "c:\\perflogs\\anvil_agent_stats.csv"
	collectorSet := "anvil_agent_win"

	// cli callbacks
	if *iniCollectorPtr {
		createCollector(collectorSet, logfilePath)
	}

	if *startCollectorPtr {
		startCollector(collectorSet)
	}

	if *runAgentPtr {
		launchServer()
	}
}
