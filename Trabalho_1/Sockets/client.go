package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"net"
	"time"
)

func generateRandomNumbers(n0 int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return n0 + r.Intn(100)
}

func main() {
	var n int
	flag.IntVar(&n, "n", 0, "number of random numbers produced")
	flag.Parse()
	conexao, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println(err)
		return
	}
	var x = 0
	for i := 0; i < n; i++ {
		x = generateRandomNumbers(x)
		str := fmt.Sprintf("%d\n", x)
		fmt.Fprintf(conexao, str)
		message, err := bufio.NewReader(conexao).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("%v Is prime: %s \n", x, message)
	}
	str := fmt.Sprintf("%d\n", 0)
	fmt.Println("Ending Execution")
	fmt.Fprintf(conexao, str)
}
