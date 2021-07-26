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
	procExists := killErr == nil

	return procExists
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputStr := scanner.Text()
		inputs := strings.Split(inputStr, ",")
		pid, _ := strconv.Atoi(inputs[0])
		signal, _ := strconv.Atoi(inputs[1])
		if processExists(pid) {
			syscall.Kill(pid, syscall.Signal(signal))
			fmt.Println("[SENDER] Signal Sended.")
		} else {
			fmt.Println("[SENDER] Couldn't find the pid", pid)
		}

	}

}
