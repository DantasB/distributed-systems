package pipe

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"syscall"

	utils "github.com/DantasB/Distributed-Systems/Utils"
)

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

		message := utils.IsPrime(n)

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
		x = utils.GenerateRandomNumbers(x)
		fmt.Fprintf(w, "%v\n", x)
	}

	fmt.Fprint(w, "0\n")
	w.Close()
}

func Pipe() {
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
