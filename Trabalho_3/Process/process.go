package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	utils "Trabalho_3/Utils"
)

func main() {
	var process_number64 uint64
	var r64 uint64
	var k int
	//Define and parse command line flags
	flag.Uint64Var(&process_number64, "pn", 0, "Size of array N")
	flag.Uint64Var(&r64, "r", 0, "Number of threads")
	flag.IntVar(&k, "k", 0, "Sleep seconds")
	flag.Parse()
	if process_number64 < 1 || r64 < 1 || k < 0 {
		fmt.Print("Incorrect flags values passed \n")
		return
	}

	var message uint32
	var r uint32 = uint32(r64)
	var process_number uint32 = uint32(process_number64)

	for i := uint32(0); i < r; i++ {
		//CoConnect to the coordinator by tcp connection
		conn, err := net.Dial("tcp", "localhost:5000")
		if err != nil {
			log.Fatalln("Couldn't connect to server")
		}

		//Send Request mesage
		binary.Write(conn, binary.BigEndian, utils.Request_message|process_number)
		//Wait and read message from server
		err = binary.Read(conn, binary.BigEndian, &message)
		if err != nil {
			log.Print("[Error] Error reading Socket\n", err)
			conn.Close()
		}
		//Checks if message was Grant as expected
		if (message & utils.Message_mask) == utils.Grant_message {
			//Was granted acess, write to file
			writeFile(fmt.Sprintf("Process number: %v\n", process_number))
			time.Sleep(time.Duration(k) * time.Second)
			//Send Release Message
			binary.Write(conn, binary.BigEndian, utils.Release_message|process_number)
		} else {
			log.Print("[Error] Message was incorrect and not granted\n")
		}

		conn.Close()
	}
}

// writeFile receives a string and returns an error.
// It will open a txt file and insert the text message on it.
// It returns the error if it ocurrs.
func writeFile(text string) error {
	//Logger needs the RDWR permission but it will just append
	file, err := os.OpenFile("resultado.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("Couldn't open file")
		return err
	}
	//Logger configuration, sets to print date with microseconds precision
	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Lmicroseconds)
	log.Printf(text)
	return file.Close()
}
