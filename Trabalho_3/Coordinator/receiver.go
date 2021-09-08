package main

import (
	"Trabalho_3/Coordinator/procqueue"
	"encoding/binary"
	"log"
	"net"
)

func receiver(pq *procqueue.ProcessQueue, abort chan<- struct{}) {
	var message uint32
	listener, err := net.Listen("tcp", "localhost:6000")
	if err != nil {
		log.Print("[Error] Error listening to socket\n")
		abort <- struct{}{}
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print("[Error] Error accepting connection\n")
		}
		err = binary.Read(conn, binary.BigEndian, message)
		if err != nil {
			conn.Close()
			log.Print("[Error] Error reading Socket\n")
		}
		if (message & message_mask) == request_message {
			processNumber := message & process_mask
			pi := procqueue.ProcessInfo{Process: processNumber, Conn: conn}
			log.Print(genMessage(message, processNumber))
			pq.Push(pi)
		} else {
			binary.Write(conn, binary.BigEndian, error_message)
			conn.Close()
			log.Print(genMessage(error_message, message&process_mask))
		}
	}
}
