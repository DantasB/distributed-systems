package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"syscall"
	"time"
)

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

func consumer(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		fmt.Printf("[SERVER] Message Received: %v\n", n)

		if n == 0 {
			fmt.Println("[SERVER] Process finished.")
			return
		}
		message := isPrime(n)

		fmt.Printf("[CLIENT] Is the value %v prime? %s \n", n, message)
	}
}

func producer(w io.WriteCloser) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("[SERVER] Write the number of the prime numbers to be generated \n")

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	x := 1
	for i := 0; i < n; i++ {
		x = generateRandomNumbers(x)
		fmt.Fprintf(w, "%v\n", x)
	}

	fmt.Fprint(w, "0\n")
	w.Close()
}

func main() {
	fds := make([]int, 2)
	err := syscall.Pipe(fds)
	if err != nil {
		fmt.Println("Pipe error:", err)
		return
	}

	r := os.NewFile(uintptr(fds[0]), "|0")
	w := os.NewFile(uintptr(fds[1]), "|1")

	id, _, _ := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)

	if id == 0 {
		producer(w)
	} else if id > 0 {
		consumer(r)
	} else {
		fmt.Println("[ERROR] Forked Failed.")
		return
	}
}
