package main

import (
	"Trabalho_3/Coordinator/procqueue"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	f, err := os.OpenFile("log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("Couldn't open file")
		os.Exit(1)
	}
	logger := log.New(f, "", log.Ldate|log.Lmicroseconds) //Concurrent Safe Logger
	abort := make(chan struct{})
	pq := procqueue.InitQueue()
	go receiver(pq, abort, logger)
	go handler(pq, logger)

	select {
	case <-abort:
		fmt.Println("Some unexpected error occurred")
		return
	case <-sigs:
		fmt.Println("Ended execution")
		return
	}

}
