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

func createCollector(collectorSet string) {
	// counters = ("\\LogicalDisk(*)\\% Free Space")

	c := exec.Command(
		"logman.exe",
		"create",
		"counter",
		collectorSet,
		"-c",
		"\"\\LogicalDisk(*)\\% Free Space\"",
		"-f",
		"csv",
		"-o",
		"c:\\perflogs\\anvil_agent_stats.csv")

	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
}

func startCollector(collectorSet string) {
	c := exec.Command(
		"logman.exe",
		"start",
		"counter",
		collectorSet)

	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err)
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
	iniCollectorPtr := flag.Bool("ini-counter", false, "Initialize logman counter")
	startCollectorPtr := flag.Bool("start-counter", false, "Start perf-agent counter")
	runAgentPtr := flag.Bool("start-agent", false, "Start a server for the agent")

	flag.Parse()

	collectorSet := "anvil_agent_win"

	// cli callbacks
	if *iniCollectorPtr {
		createCollector(collectorSet)
	}

	if *startCollectorPtr {
		startCollector(collectorSet)
	}

	if *runAgentPtr {
		launchServer()
	}
}
