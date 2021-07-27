package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"
)

func processExists(pid int) bool {
	killErr := syscall.Kill(pid, syscall.Signal(0))
	return killErr == nil
}

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
