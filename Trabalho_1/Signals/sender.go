package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputStr := scanner.Text()
		inputs := strings.Split(inputStr, ",")
		pid, _ := strconv.Atoi(inputs[0])
		signal, _ := strconv.Atoi(inputs[1])

		syscall.Kill(pid, syscall.Signal(signal))

		fmt.Println("[SENDER] Signal Sended.")
	}

}
