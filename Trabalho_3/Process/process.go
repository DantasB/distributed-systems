package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"

	utils "Trabalho_3/Utils"
)

func main() {
	var process_number uint32 = 1
	var k uint32 = 3
	var message uint32

	for i := uint32(0); i < k; i++ {
		conn, err := net.Dial("tcp", "localhost:6000")
		if err != nil {
			log.Fatalln("Counden't connect to server")
		}
		binary.Write(conn, binary.BigEndian, utils.Request_message|process_number)
		err = binary.Read(conn, binary.BigEndian, &message)
		if err != nil {
			log.Print("[Error] Error reading Socket\n", err)
			conn.Close()
		}
		if (message & utils.Message_mask) == utils.Grant_message {
			writeFile()
			binary.Write(conn, binary.BigEndian, utils.Release_message|process_number)
		} else {
			log.Print("[Error] Message was incorrect and not granted\n")
		}
		conn.Close()
	}
}

func writeFile() {
	fmt.Printf("Writing to file \n")
}
