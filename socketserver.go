package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"strings" // only needed below for sample processing
)

// Performance Monitor
// relog 'C:\Path' -o text.csv -f csv

func createCounter() {

}

func startCounter() {

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
	iniCounterPtr := flag.Bool("ini-counter", false, "Initialize logman counter")
	startCounterPtr := flag.Bool("start-counter", false, "Start perf-agent counter")
	runAgentPtr := flag.Bool("start-agent", false, "Start a server for the agent")

	flag.Parse()

	// cli callbacks
	if *iniCounterPtr {
		createCounter()
	}

	if *startCounterPtr {
		startCounter()
	}

	if *runAgentPtr {
		launchServer()
	}
}
