package main

import (
	"context"
	"flag"
	"fmt"
	"math"
	"math/rand"
	"time"

	semaphore "golang.org/x/sync/semaphore"
)

//consumer limit
var m = 100000

//memory instantiation
var memory []int

//seed of random number
var seed = rand.New(rand.NewSource(time.Now().UnixNano()))

//semaphores
var empty *semaphore.Weighted
var full *semaphore.Weighted
var mutex *semaphore.Weighted
var loopControl *semaphore.Weighted

// generateRandomNumber receives nothing and returns a integer.
// It will use a seed that generates a number from 1 to 10^7.
// It returns the integer random number.
func generateRandomNumber() int {
	return seed.Intn(int(math.Pow(10, 7))) + 1
}

// createArrayWithZeros receives n and returns a vector of zeroes.
// It will instantiate a vector with size n.
// It returns the array containing only zeros.
func createArrayWithZeros(n int) []int {
	return make([]int, n)
}

// GetSquareRoot receives an integer number and returns this square root.
// It's necessary to cast the integer to float64 because of the sqrt function.
// Ceil the obtained square root because the output is float64.
// Convert again to integer because of the output of this function.
// It returns the square root of a number.
func getSquareRoot(number int) int {
	return int(math.Ceil(math.Sqrt(float64(number))))
}

// isPrime receives an integer number and returns a string.
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

// getFreePosition receives nothing and returns a integer.
// It will iterate over the array.
// Check if the value of the index is 0.
// It returns -1 if there's no empty space.
// or returns the index if the array in that index is equal 0.
func getFreePosition() int {
	for i := 0; i < len(memory); i++ {
		if memory[i] == 0 {
			return i
		}
	}
	return -1
}

// getFirstFullPosition receives nothing and returns a integer.
// It will iterate over the array.
// Check if the value of the index is different from 0.
// It returns the index if the array in that index is different from 0.
// or returns -1 if all values of the array is equal 0.
func getFirstFullPosition() int {
	for i := 0; i < len(memory); i++ {
		if memory[i] != 0 {
			return i
		}
	}
	return -1
}

// consumes receives an index of a global array.
// checks if this index contains a prime value.
// fills this index position with 0
func consumes() {
	var value = getFirstFullPosition()
	fmt.Printf("Is Value %d Prime? %s\n", memory[value], isPrime(memory[value]))
	memory[value] = 0
}

// produces receives an index of a global array and fills it with a random number.
func produces() {
	memory[getFreePosition()] = generateRandomNumber()
}

// producer will create a context and loop forever producing.
func producer() {
	ctx := context.Background()
	for {
		empty.Acquire(ctx, 1)
		mutex.Acquire(ctx, 1)
		produces()
		mutex.Release(1)
		full.Release(1)
	}

}

// consumer will receive a finished channel as parameter and return none.
// it will create a context
// the first steps are to acquire the loopControl mutex and decrease the m value (limit)
// after that it will release
// at the end of the loop, it will send true to the channel (finished consuming)
func consumer(finished chan bool) {
	ctx := context.Background()
	for {
		loopControl.Acquire(ctx, 1)
		if m == 0 {
			loopControl.Release(1)
			break
		}
		m--
		loopControl.Release(1)

		full.Acquire(ctx, 1)
		mutex.Acquire(ctx, 1)
		consumes()
		mutex.Release(1)
		empty.Release(1)
	}
	finished <- true
}

//setFullToZero acquires all semaphores resources to start full as zero.
func setFullToZero(n int) {
	ctx := context.Background()
	for i := 0; i < n; i++ {
		full.Acquire(ctx, 1)
	}
}

func main() {
	var np int
	var nc int
	var n int
	flag.IntVar(&np, "np", 0, "Number of Producer Threads")
	flag.IntVar(&nc, "nc", 0, "Number of Consumer Threads")
	flag.IntVar(&n, "n", 0, "Shared Memory Size")
	flag.Parse()

	if n < 1 || nc < 1 || np < 1 {
		fmt.Print("Incorrect flags values passed \n")
		return
	}

	finished := make(chan bool, nc)
	memory = createArrayWithZeros(n)
	full = semaphore.NewWeighted(int64(n))
	setFullToZero(n)
	empty = semaphore.NewWeighted(int64(n))
	mutex = semaphore.NewWeighted(1)
	loopControl = semaphore.NewWeighted(1)

	start := time.Now()
	for i := 0; i < nc; i++ {
		go consumer(finished)
	}

	for i := 0; i < nc; i++ {
		go producer()
	}

	var consumersEnded int
	for consumersEnded != nc {
		<-finished
		consumersEnded++
	}

	duration := time.Since(start)
	fmt.Printf("%v", duration.Seconds())
}
