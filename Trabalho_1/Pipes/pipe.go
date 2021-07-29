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

// GenerateRandomNumbers receives a integer x and returns a random number.
// It will use a seed that generates a number from 0 to 99.
// Sum the x + 1 with this random number.
// It returns the a integer random number.
func generateRandomNumbers(x int) int {
	return x + seed.Intn(100) + 1
}

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

// Consumer receives a io.Reader and has no return.
// It will instantiate a new scanner with the reader.
// Waits for every element scanned in the buffer.
// Converts the scanner content to integer.
// Verifies if the received content is 0, if yes it exits the function.
// Check if the content is a prime number.
// Prints the message and the value.
// It has no return.
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

// Producer receives a io.WriteCloser and has no return.
// It will instantiate a new scanner with the os.Stdin.
// Waits for the user input.
// Reads and convert the user input to integer.
// Iterates from 0 to user input - 1.
// Generate a Random Number from the previous x value
// Writes it to the w Writer.
// After the end of the loop it writes 0 to the w Writer and close the Writer.
// It has no return.
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
