package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"net"
	"time"
)

var seed = rand.New(rand.NewSource(time.Now().UnixNano()))

// GenerateRandomNumbers receives an integer x and returns a random number.
// It will use a seed that generates a number from 0 to 99.
// Sum the x + 1 with this random number.
// It returns the a integer random number.
func generateRandomNumbers(x int) int {
	return x + seed.Intn(100) + 1
}

func main() {
	var n int
	flag.IntVar(&n, "n", 0, "Number of random numbers produced")
	flag.Parse()

	conexao, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println(err)
		return
	}

	var x = 1
	for i := 0; i < n; i++ {
		x = generateRandomNumbers(x)
		str := fmt.Sprintf("%d\n", x)
		fmt.Fprintf(conexao, str)
		message, err := bufio.NewReader(conexao).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("[CLIENT] Is the value %v prime? %s \n", x, message)
	}

	str := fmt.Sprintf("%d\n", 0)
	fmt.Println("[CLIENT] Process finished.")
	fmt.Fprintf(conexao, str)
}
