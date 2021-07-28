package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"
)

var r, w *os.File

func generateRandomNumbers(x int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return x + r.Intn(100)
}

func isPrime(number int) string {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return "false"
		}
	}
	return "true"
}

func consumer() {
	buffer := new(strings.Builder)
	io.Copy(buffer, r)

	n, _ := strconv.Atoi(buffer.String())
	fmt.Print("[SERVER] Message Received : ", n)

	if n == 0 {
		fmt.Println("[SERVER] Process finished.")
		return
	}
	message := isPrime(n)

	fmt.Printf("[CLIENT] Is the value %v prime? %s \n", n, message)
}

func producer() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("[SERVER] Write the number of the prime numbers to be generated \n")

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	var x = 1
	for i := 0; i < n; i++ {
		x = generateRandomNumbers(x)
		fmt.Printf("%d \n", x)
		fmt.Fprint(w, x)
	}

	fmt.Fprint(w, 0)

	w.Close()
}

func main() {

	r, w, _ = os.Pipe()

	id, _, _ := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)

	if id == 0 {
		consumer()
	} else if id > 0 {
		producer()
	} else {
		fmt.Println("[ERROR] Forked Failed.")
	}
}
