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

// To this process all the results are written in the same file.
var processLogPath string = "../Process/resultado.txt"
var coordinatorLogPath string = "../Coordinator/log"

// ParseLogTime receives datetime string and returns a time.Time and an error.
// It will try to parse a string containing the datetime with a layout.
// If its not possible to parse, it will return a time and the error.
// If its possible it will return the parsed time and nil as the error.
func ParseLogTime(datetime string) (time.Time, error) {
	layout := "2006/01/02 15:04:05.000000"

	parsedTime, err := time.Parse(layout, datetime)
	if err != nil {
		fmt.Println("Error while parsing date:", datetime)
		return time.Unix(0, 0), err
	}

	return parsedTime, nil
}

// InvalidLogTime receives a line, a string separator, the startTime time and returns a boolean.
// It will split the line with the separator.
// It will try to parse the datetime string.
// If it could parse, it will check if this time ocurred before the start time.
// If not it will return true, indicating that an error ocurred.
// It will set the start time as the parsedTime.
// It will returns a boolean. True if it is valid, else false.
func InvalidLogTime(line string, separator string, startTime *time.Time) bool {
	splittedLine := strings.Split(line, separator)
	parsedTime, err := ParseLogTime(splittedLine[0])
	if err != nil {
		return true
	}

	if (*startTime).After(parsedTime) {
		return true
	}

	*startTime = parsedTime

	return false
}

// IsValidProcessLogs receives an int n, an int r and returns a boolean.
// It will open the process log path and reads every line.
// It will check if the number of lines is equal n*r.
// It will also check for every line if the time is valid or not.
// It will returns a boolean. If it's valid returns true, else false.
func IsValidProcessLogs(n int, r int) bool {
	numberOfLines := 0

	startTime := time.Unix(0, 0)

	file, err := os.OpenFile(processLogPath, os.O_RDONLY, 0777)
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

// GetProcessNumber receives a string line and returns an int and an error.
// It will get the last ocurrence of a ':' string and gets the number after that.
// It will convert it to integer and returns the number and the error.
func GetProcessNumber(line string) (int, error) {
	return strconv.Atoi(line[strings.LastIndex(line, ":")+2:])
}

// GrantIsIncorrect receives a processOcurrence, lastGrantProcess and a process number and returns a bool.
// It will check the len of the processOcurrence queue.
// It will check if the first element of the queue is equal the process number.
// It will also check if the last grant process if different from -1.
// It will remove the first element of the queue and redefine the last grant process.
// It returns true if any error ocurred, false else.
func GrantIsIncorrect(processOcurrence *[]int, lastGrantProcess *int, processNumber int) bool {
	if len((*processOcurrence)) == 0 {
		return true
	}

	firstElement := (*processOcurrence)[0]
	if firstElement != processNumber || (*lastGrantProcess) != -1 {
		return true
	}

	*processOcurrence = (*processOcurrence)[1:]
	*lastGrantProcess = processNumber

	return false
}

// ReleaseIsIncorrect receives a processNumber, lastGrantProcess and returns a bool.
// It will check if the process number is equals -1.
// It will also check if the processNumber is different from the last grant process.
// It will redefine the last grant process to -1.
// It returns true if any error ocurred, false else.
func ReleaseIsIncorrect(processNumber int, lastGrantProcess *int) bool {
	if processNumber == -1 || processNumber != (*lastGrantProcess) {
		return true
	}

	*lastGrantProcess = -1
	return false
}

// IsValidCoordinatorLogs receives nothing and returns a boolean.
// It will read the coordinator log path and reads every line.
// It will check the message type (Request, Grant or Release) and treat it.
// For each Request message it will append to the processOcurrence queue.
// For each Grant message it will check if the grant is incorrect.
// For each Release message it will check if the release is incorrect.
// It will return a boolean true if no errors ocurred, false else.
func IsValidCoordinatorLogs() bool {
	var processOcurrence []int

	file, err := os.OpenFile(coordinatorLogPath, os.O_RDONLY, 0777)
	if err != nil {
		fmt.Println("Couldn't open file")
		return false
	}

	scanner := bufio.NewScanner(file)
	lastGrantProcess := -1
	for scanner.Scan() {
		line := scanner.Text()
		processNumber, _ := GetProcessNumber(line)
		switch {
		case strings.Contains(line, "Request"):
			processOcurrence = append(processOcurrence, processNumber)

		case strings.Contains(line, "Grant"):
			if GrantIsIncorrect(&processOcurrence, &lastGrantProcess, processNumber) {
				return false
			}

		case strings.Contains(line, "Release"):
			if ReleaseIsIncorrect(processNumber, &lastGrantProcess) {
				return false
			}
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
	if !IsValidProcessLogs(n, r) {
		fmt.Print("[Check failed] Invalid Process log\n")
	} else if !IsValidCoordinatorLogs() {
		fmt.Print("[Check failed] Invalid Coordinator log\n")
	} else {
		fmt.Print("Check successful \n")
	}
}
