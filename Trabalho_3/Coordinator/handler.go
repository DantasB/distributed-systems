package main

import (
	"Trabalho_3/Coordinator/procqueue"
	"encoding/binary"
	"log"

	utils "Trabalho_3/Utils"
)

//Pops the first process of the queue and grant acess to critical region until the program exits
func handler(pq *procqueue.ProcessQueue, logger *log.Logger) {
	var message uint32
	for {
		pi := pq.Pop()
		// Send the grant message to process that was the first in the queue
		err := binary.Write(pi.Conn, binary.BigEndian, utils.Grant_message)
		if err != nil {
			log.Print("[Error] Error writing to Socket\n")
			pi.Conn.Close()
			continue
		}
		//Logs the message sent
		logger.Print(utils.GenMessage(utils.Grant_message, pi.Process))
		//Wait and read the message from the process
		err = binary.Read(pi.Conn, binary.BigEndian, &message)
		if err != nil {
			log.Print("[Error] Error reading Socket\n", err)
			pi.Conn.Close()
			continue
		}
		//Logs the message received
		logger.Print(utils.GenMessage(message, pi.Process))
		//Checks if the message was received as expected
		if (message & utils.Message_mask) != utils.Release_message {
			binary.Write(pi.Conn, binary.BigEndian, utils.Error_message)
		}
		pi.Conn.Close()
	}
}
