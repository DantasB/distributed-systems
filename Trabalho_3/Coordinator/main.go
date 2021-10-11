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

// amountParsing receives a string and returns a tuple with an uint32 and an error.
// It will parse a input string and get the process id if the function is named amount.
// It returns the pid as uint32 and the error message if it exists.
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

// terminal receives a ProcessQueue and a channel, it returns nothing.
// It will parse the os.Stding input and initializes the terminal mode.
// The therminal mode receives a input and prints the informations about the coordinator.
// Also the terminal will print the informations about the process.
// It returns nothing.
func terminal(pq *procqueue.ProcessQueue, abort chan struct{}) {
	reader := bufio.NewReader(os.Stdin)
	printHelpMessage := func() {
		fmt.Printf("           Kill. Kills the process.\n")
		fmt.Printf("           Queue. Returns the actual queue.\n")
		fmt.Printf("           Amount, process_number. Returns the amount of grants to this process number\n")
	}
	fmt.Printf("[TERMINAL] Write one of the three modes:\n")
	printHelpMessage()
	for {
		select {
		case <-abort:
			fmt.Println("\nSome unexpected error occurred")
			return
		default:
			fmt.Printf("> ")
			n, _ := reader.ReadString('\n')
			input := strings.ToLower(strings.Trim(n, "\n"))
			if strings.Trim(input, " ") == "kill" {
				os.Exit(0)
			} else if strings.Trim(input, " ") == "queue" {
				fmt.Printf("[TERMINAL] This is the actual Queue: %v", pq.Print())
			} else {
				process_number, error := amountParsing(input)
				if error == nil {
					fmt.Printf("[TERMINAL] This is the amount of time that the process has been granted access by the coordinator %v: %v\n", process_number, pq.Count(process_number))
				} else {
					fmt.Printf("%v", error)
					fmt.Printf("[TERMINAL] Couldn't recognize this instruction. Try to use the following options:\n")
					printHelpMessage()
				}
			}
		}
	}
}

func main() {
	fmt.Println("Starting Coordinator. Waiting for messages...\n")
	//Config and initializes the log
	f, err := os.OpenFile("log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("Couldn't open file")
		os.Exit(1)
	}
	logger := log.New(f, "", log.Ldate|log.Lmicroseconds)

	abort := make(chan struct{})
	//Initializes the thread-safe process queue
	pq := procqueue.InitQueue()

	//Start the necessary go routines
	go receiver(pq, abort, logger)
	go handler(pq, logger)
	terminal(pq, abort)
}
