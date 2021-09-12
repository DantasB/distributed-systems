package main

import (
	"Trabalho_3/Coordinator/procqueue"
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func amountParsing(input string) (uint32, error) {
	inputs := strings.Split(strings.Trim(input, " "), ",")
	if len(inputs) != 2 {
		return 0, errors.New("Incorrect Input")
	}

	function := strings.Trim(inputs[0], " ")
	if function != "amount" {
		return 0, errors.New("Invalid Function")
	}

	pid, _ := strconv.Atoi(strings.Trim(inputs[1], " "))
	if pid <= 0 {
		return 0, errors.New("Invalid PID")
	}

	return uint32(pid), nil
}

func terminal(pq *procqueue.ProcessQueue, abort chan struct{}) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("[TERMINAL] Write one of the three modes:\n")
	fmt.Printf("           Kill. Kills the process.\n")
	fmt.Printf("           Queue. Returns the actual queue.\n")
	fmt.Printf("           Amount, process_number. Returns the amount of this process number\n")

	for {
		select {
		case <-abort:
			fmt.Println("\nSome unexpected error occurred")
			return
		default:
			n, _ := reader.ReadString('\n')
			input := strings.ToLower(strings.Trim(n, "\n"))
			if strings.Trim(input, " ") == "kill" {
				os.Exit(1)
			} else if strings.Trim(input, " ") == "queue" {
				fmt.Printf("[TERMINAL] This is the actual Queue: %v\n", pq.Print())
			} else {
				process_number, error := amountParsing(input)
				if error == nil {
					fmt.Printf("[TERMINAL] This is the amount of time that the process has been granted access by the coordinator %v: %v\n", process_number, pq.Count(process_number))
				} else {
					fmt.Printf("%v", error)
					fmt.Printf("[TERMINAL] Couldn't recognize this instruction. Try to use the following options:\n")
					fmt.Printf("           Kill. Kills the process.\n")
					fmt.Printf("           Queue. Returns the actual queue.\n")
					fmt.Printf("           Amount, process_number. Returns the amount of this process number\n")
				}
			}
		}
	}
}

func main() {
	fmt.Println("Starting Coordinator. Waiting for messages...\n")

	f, err := os.OpenFile("log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("Couldn't open file")
		os.Exit(1)
	}
	logger := log.New(f, "", log.Ldate|log.Lmicroseconds)
	abort := make(chan struct{})
	pq := procqueue.InitQueue()
	go receiver(pq, abort, logger)
	go handler(pq, logger)
	terminal(pq, abort)
}
