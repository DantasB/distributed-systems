package sockets

import (
	"bufio"
	"fmt"
	"net"

	utils "github.com/DantasB/distributed-systems/Utils"
)

func SocketServer() {
	fmt.Println("[SERVER] Awaiting connection...")

	//Open TCP server on port 8081
	server, err := net.Listen("tcp", ":8081")
	//Checks if error is not null and end execution if it is
	if err != nil {
		fmt.Println(err)
		return
	}

	//Accepts a new connection
	conn, err := server.Accept()
	//Checks if error is not null and end execution if it is
	if err != nil {
		fmt.Println(err)
		return
	}

	defer server.Close()

	fmt.Println("[SERVER] Connection Accepted. Waiting for a message.")
	//Infinite loop receiving numbers on the connection and ending if the number is equal to 0
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print("[SERVER] Message Received : ", message)

		var val int
		_, err = fmt.Sscanf(message, "%d\n", &val)
		if err != nil {
			fmt.Println(err)
			return
		}

		if val == 0 {
			fmt.Println("[SERVER] Process finished.")
			return
		}

		novamensagem := fmt.Sprintf("%v", utils.IsPrime(val))
		conn.Write([]byte(novamensagem + "\n"))
	}

}
