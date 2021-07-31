package main

import (
	"bufio"
	"fmt"
	"math"
	"net"
)

// GetSquareRoot receives an integer number and returns this square root.
// It's necessary to cast the integer to float64 because of the sqrt function.
// Ceil the obtained square root because the output is float64.
// Convert again to integer because of the output of this function.
// It returns the square root of a number.
func getSquareRoot(number int) int {
	return int(math.Ceil(math.Sqrt(float64(number))))
}

// IsPrime receives an integer number and returns a string.
// It will iterate over 2 to the square root of the number - 1.
// Check if the number is divisible by the i.
// It returns the string false or true.
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

	fmt.Println("[SERVER] Connection Accepted.")
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

		novamensagem := fmt.Sprintf("%v", isPrime(val))
		conn.Write([]byte(novamensagem + "\n"))
	}

}
