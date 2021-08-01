package sockets

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"

	utils "github.com/DantasB/Distributed-Systems/Utils"
)

func SocketClient() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("[Client] Write the amount of numbers to be generated \n")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	//Connets to tcp server
	conexao, err := net.Dial("tcp", "127.0.0.1:8081")
	//Checks Error and end execution if it not null
	if err != nil {
		fmt.Print(err)
		return
	}

	var x = 1
	//Gets n random numbers and sends do TCP server where n is the flag passed by the user
	for i := 0; i < n; i++ {
		x = utils.GenerateRandomNumbers(x)
		str := fmt.Sprintf("%d\n", x)
		fmt.Fprintf(conexao, str)
		message, err := bufio.NewReader(conexao).ReadString('\n')
		if err != nil {
			fmt.Print(err)
			return
		}

		fmt.Printf("[CLIENT] Is the value %v prime? %s", x, message)
	}

	//Send 0 to the tcp server to end execution and ends
	str := fmt.Sprintf("%d\n", 0)
	fmt.Println("[CLIENT] Process finished.")
	fmt.Fprintf(conexao, str)
}
