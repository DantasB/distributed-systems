package main

import (
	"Trabalho_3/Coordinator/procqueue"
	"encoding/binary"
	"log"

	utils "Trabalho_3/Utils"
)

func handler(pq *procqueue.ProcessQueue, logger *log.Logger) {
	var message uint32
	for {
		pi := pq.Pop()
		err := binary.Write(pi.Conn, binary.BigEndian, utils.Grant_message)
		if err != nil {
			log.Print("[Error] Error writing to Socket\n")
			pi.Conn.Close()
			continue
		}
		logger.Print(utils.GenMessage(utils.Grant_message, pi.Process))
		err = binary.Read(pi.Conn, binary.BigEndian, &message)
		if err != nil {
			log.Print("[Error] Error reading Socket\n", err)
			pi.Conn.Close()
			continue
		}
		logger.Print(utils.GenMessage(message, pi.Process))
		if (message & utils.Message_mask) != utils.Release_message {
			binary.Write(pi.Conn, binary.BigEndian, utils.Error_message)
		}
		pi.Conn.Close()
	}
}
