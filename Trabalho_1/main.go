package main

import (
	"flag"
	"fmt"

	pipe "github.com/DantasB/Distributed-Systems/Pipes"
	signals "github.com/DantasB/Distributed-Systems/Signals"
	sockets "github.com/DantasB/Distributed-Systems/Sockets"
)

func main() {
	var programName string
	//Collect the flags passed by the user
	flag.StringVar(&programName, "program_name", "", "Name of the program to be executed")
	flag.Parse()
	switch programName {
	case "socket_client":
		sockets.SocketClient()
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
