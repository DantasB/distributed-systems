package main

import (
	"Trabalho_3/Coordinator/procqueue"
	"encoding/binary"
	"log"
)

func handler(pq *procqueue.ProcessQueue) {
	var message uint32
	for {
		pi := pq.Pop()
		err := binary.Write(pi.Conn, binary.BigEndian, grant_message)
		if err != nil {
			log.Print("[Error] Error writing to Socket\n")
			pi.Conn.Close()
			continue
		}
		log.Print(genMessage(grant_message, pi.Process))
		err = binary.Read(pi.Conn, binary.BigEndian, message)
		if err != nil {
			log.Print("[Error] Error reading Socket\n")
			pi.Conn.Close()
			continue
		}
		log.Print(genMessage(message, pi.Process))
		if (message & message_mask) != release_message {
			binary.Write(pi.Conn, binary.BigEndian, error_message)
		}
		pi.Conn.Close()
	}
}
