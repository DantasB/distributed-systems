package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"time"
)

//consumer limit
var m = 100000

//seed of random number
var seed = rand.New(rand.NewSource(time.Now().UnixNano()))

// generateRandomNumber receives nothing and returns a integer.
// It will use a seed that generates a number from 1 to 10^7.
// It returns the integer random number.
func generateRandomNumber() int {
	return seed.Intn(int(math.Pow(10, 7))) + 1
}

// createArrayWithZeros receives nothing and returns a vector.
// It will instantiate a vector with size m.
// It will iterate over the array and set the value of 0 to 2 positions.
// It returns the array containing only zeros.
func createArrayWithZeros() []int {
	memory := make([]int, m)
	memory[0] = 0
	for i := 1; i < len(memory); i *= 2 {
		copy(memory[i:], memory[:i])
	}

	return memory
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

// getFreePosition receives an array of integers and returns a integer.
// It will iterate over the array.
// Check if the value of the index is 0.
// It returns -1 if there's no empty space.
// or returns the index if the array in that index is equal 0.
func getFreePosition(memory []int) int {
	for i := 0; i < len(memory); i++ {
		if memory[i] == 0 {
			return i
		}
	}
	return -1
}

// getFirstFullPosition receives an array of integers and returns a integer.
// It will iterate over the array.
// Check if the value of the index is different from 0.
// It returns the index if the array in that index is different from 0.
// or returns -1 if all values of the array is equal 0.
func getFirstFullPosition(memory []int) int {
	for i := 0; i < len(memory); i++ {
		if memory[i] != 0 {
			return i
		}
	}
	return -1
}

// isEmpty receives an array of integers and returns a boolean.
// It will check if there's any array index with value different from 0.
// It returns false if yes and true if the array is full of zeros.
func isEmpty(memory []int) bool {
	if getFirstFullPosition(memory) == -1 {
		return true
	}

	return false
}

// isFull receives an array of integers and returns a boolean.
// It will check if there's any array index with value equals 0.
// It returns false if yes and true if the array contains all values different from 0.
func isFull(memory []int) bool {
	if getFreePosition(memory) == -1 {
		return true
	}

	return false
}

func producer(vec []int) {

}

func consumer(vec []int) {

}

func main() {
	var np int
	var nc int
	var n int
	flag.IntVar(&nc, "nc", 0, "Number of Consumer Threads")
	flag.IntVar(&np, "kp", 0, "Number of Producer Threads")
	flag.IntVar(&n, "n", 0, "Shared Memory Size")
	flag.Parse()
	if n < 1 || nc < 1 || np < 1 {
		fmt.Print("Incorrect flags values passed \n")
		return
	}

}
