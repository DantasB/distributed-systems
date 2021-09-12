package main

import (
	"Trabalho_3/Coordinator/procqueue"
	"encoding/binary"
	"log"
	"net"

	utils "Trabalho_3/Utils"
)

func receiver(pq *procqueue.ProcessQueue, abort chan<- struct{}, logger *log.Logger) {
	var message uint32
	listener, err := net.Listen("tcp", "localhost:6000")
	if err != nil {
		logger.Print("[Error] Error listening to socket\n")
		abort <- struct{}{}
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Print("[Error] Error accepting connection\n")
		}

		err = binary.Read(conn, binary.BigEndian, &message)
		if err != nil {
			conn.Close()
			logger.Print("[Error] Error reading Socket\n", err)
		}

		if (message & utils.Message_mask) == utils.Request_message {
			processNumber := message & utils.Process_mask
			pi := procqueue.ProcessInfo{Process: processNumber, Conn: conn}
			logger.Print(utils.GenMessage(message, processNumber))
			pq.Push(pi)
		} else {
			binary.Write(conn, binary.BigEndian, utils.Error_message)
			conn.Close()
			logger.Print(utils.GenMessage(utils.Error_message, message&utils.Process_mask))
		}
	}
}
