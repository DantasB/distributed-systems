package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"strconv"
	"syscall"
	"time"
)

var seed = rand.New(rand.NewSource(time.Now().UnixNano()))

func generateRandomNumbers(x int) int {
	return x + seed.Intn(100) + 1
}

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

func consumer(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())

		fmt.Printf("[CONSUMER] Message Received: %v\n", n)

		if n == 0 {
			fmt.Println("[CONSUMER] Process finished.")
			return
		}

		message := isPrime(n)

		fmt.Printf("[CONSUMER] Is the value %v prime? %s \n", n, message)
	}
}

func producer(w io.WriteCloser) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("[PRODUCER] Write the number of the prime numbers to be generated \n")

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
	pipe := make([]int, 2)
	syscall.Pipe(pipe)

	r := os.NewFile(uintptr(pipe[0]), "consumer")
	w := os.NewFile(uintptr(pipe[1]), "producer")

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
