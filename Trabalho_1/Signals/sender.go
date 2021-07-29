package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"
)

// ProcessExists receives an integer pid and returns true if this process exists.
// It makes a Kill syscall for the pid with a signal 0 (Check access to pid).
// If it's nil that means that the process exists.
// It returns true if the process exists and false if it doesn't.
func processExists(pid int) bool {
	killErr := syscall.Kill(pid, syscall.Signal(0))
	return killErr == nil
}

// InputParsing receives an input string and returns the parsed input as tuple.
// It splits the input containing the comma and assigns it to a inputs variable.
// Trim the every element of the inputs with a empty space and converts to a integer.
// It returns a tuple containing the pid and the signal inputed.
func inputParsing(input string) (int, int) {
	inputs := strings.Split(input, ",")
	pid, _ := strconv.Atoi(strings.Trim(inputs[0], " "))
	signal, _ := strconv.Atoi(strings.Trim(inputs[1], " "))

	return pid, signal
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		pid, signal := inputParsing(scanner.Text())
		if processExists(pid) {
			syscall.Kill(pid, syscall.Signal(signal))
			fmt.Println("[SENDER] Signal Sended.")
		} else {
			fmt.Println("[SENDER] Couldn't find the pid", pid)
		}

	}

}
