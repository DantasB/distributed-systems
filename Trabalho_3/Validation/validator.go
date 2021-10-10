package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var processLogPath string = "../Process/resultado.txt"
var coordinatorLogPath string = "../Coordinator/log"

func ParseLogTime(datetime string) time.Time {
	layout := "2006/01/02 15:04:05.000000"

	parsedTime, err := time.Parse(layout, datetime)
	if err != nil {
		fmt.Println("Error while parsing date:", err)
	}

	return parsedTime
}

func InvalidLogTime(line string, splitString string, startTime *time.Time) bool {
	datetime := strings.Split(line, splitString)
	parsedTime := ParseLogTime(datetime[0])
	if (*startTime).After(parsedTime) {
		return true
	}

	*startTime = parsedTime

	return false
}

func IsValidProcessLogs(path string, n int, r int) bool {
	numberOfLines := 0

	startTime := time.Unix(0, 0)

	file, err := os.OpenFile(path, os.O_RDONLY, 0777)
	if err != nil {
		fmt.Println("Couldn't open file")
		return false
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numberOfLines++
		if InvalidLogTime(scanner.Text(), " Process", &startTime) {
			return false
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	if numberOfLines != n*r {
		return false
	}

	return true
}

func IsValidCoordinatorLogs(path string) bool {
	var processOcurrence []int

	file, err := os.OpenFile(path, os.O_RDONLY, 0777)
	if err != nil {
		fmt.Println("Couldn't open file")
		return false
	}

	scanner := bufio.NewScanner(file)
	lastGrantProcess := -1
	for scanner.Scan() {
		line := scanner.Text()
		processNumber, _ := strconv.Atoi(line[strings.LastIndex(line, ":")+2:])
		switch {
		case strings.Contains(line, "Request"):
			processOcurrence = append(processOcurrence, processNumber)

		case strings.Contains(line, "Grant"):
			if len(processOcurrence) == 0 {
				return false
			}

			firstElement := processOcurrence[0]
			if firstElement != processNumber || lastGrantProcess != -1 {
				return false
			}

			processOcurrence = processOcurrence[1:]
			lastGrantProcess = processNumber

		case strings.Contains(line, "Release"):
			if processNumber == -1 || processNumber != lastGrantProcess {
				return false
			}

			lastGrantProcess = -1
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return true
}

func main() {
	var r, n int
	flag.IntVar(&r, "r", 0, "Number of threads")
	flag.IntVar(&n, "n", 0, "Number of processes")
	flag.Parse()
	if r < 1 || n < 1 {
		fmt.Print("Incorrect flags values passed \n")
		return
	}
	if !IsValidProcessLogs(processLogPath, n, r) {
		fmt.Print("[Check failed] Invalid Process log\n")
	} else if !IsValidCoordinatorLogs(coordinatorLogPath) {
		fmt.Print("[Check failed] Invalid Coordinator log\n")
	} else {
		fmt.Print("Check successful \n")
	}
}
