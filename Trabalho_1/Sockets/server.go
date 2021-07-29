package main

import (
	"bufio"
	"fmt"
	"math"
	"net"
	"os"
)

func getSquareRoot(number int) int {
	return int(math.Ceil(math.Sqrt(float64(number))))
}

func isPrime(number int) string {
	for i := 2; i < getSquareRoot(number); i++ {
		if number%i == 0 {
			return "false"
		}
	}
	return "true"
}

func main() {
	fmt.Println("[SERVER] Awaiting connection...")

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

	fmt.Println("[SERVER] Connection Accepted.")
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

		novamensagem := fmt.Sprintf("%v", isPrime(val))
		conn.Write([]byte(novamensagem + "\n"))
	}

}
