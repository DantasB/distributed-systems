package main

import (
	"Trabalho_3/Coordinator/procqueue"
	"fmt"
)

func main() {
	abort := make(chan struct{})
	pq := procqueue.InitQueue()
	go receiver(pq, abort)
	go handler(pq)
	<-abort
	fmt.Println("Execution ended")
}
