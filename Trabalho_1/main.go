package main

import (
	"flag"
	"fmt"

	pipe "github.com/DantasB/Distributed-Systems/Pipes"
	signals "github.com/DantasB/Distributed-Systems/Signals"
	sockets "github.com/DantasB/Distributed-Systems/Sockets"
)

func main() {
	var n int
	var programName string
	flag.StringVar(&programName, "program_name", "", "Name of the program to be executed")
	flag.IntVar(&n, "n", 0, "Number of random numbers produced")
	flag.Parse()
	switch programName {
	case "socket_client":
		sockets.SocketClient(n)
	case "socket_server":
		sockets.SocketServer()
	case "signal_rec":
		signals.SignalReceiver()
	case "signal_sen":
		signals.SignalSender()
	case "pipe":
		pipe.Pipe()
	default:
		fmt.Println("Incorrect program name")
	}
	return
}
