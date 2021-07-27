package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("Server awaiting connection...")
	server, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	conn, err := server.Accept()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer server.Close()

	fmt.Println("Connection Accepted...")
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print("Message Received : ", message)
		var val int
		_, err = fmt.Sscanf(message, "%d\n", &val)
		if err != nil {
			fmt.Println(err)
			return
		}
		if val == 0 {
			fmt.Println("Reciveied 0 \nEnding Execution")
			return
		}
		is_prime := true
		for i := 2; i < val; i++ {
			if val%i == 0 {
				is_prime = false
			}
		}
		novamensagem := fmt.Sprintf("%v", is_prime)
		conn.Write([]byte(novamensagem + "\n"))
	}

}
