package utils

import (
	"math"
	"math/rand"
	"time"
)

// GetSquareRoot receives an integer number and returns this square root.
// It's necessary to cast the integer to float64 because of the sqrt function.
// Ceil the obtained square root because the output is float64.
// Convert again to integer because of the output of this function.
// It returns the square root of a number.
func getSquareRoot(number int) int {
	return int(math.Ceil(math.Sqrt(float64(number))))
}

// IsPrime receives an integer number and returns a string.
// It will iterate over 2 to the square root of the number - 1.
// Check if the number is divisible by the i.
// It returns the string false or true.
func IsPrime(number int) string {
	for i := 2; i < getSquareRoot(number); i++ {
		if number%i == 0 {
			return "false"
		}
	}
	return "true"
}

var seed = rand.New(rand.NewSource(time.Now().UnixNano()))

// GenerateRandomNumbers receives an integer x and returns a random number.
// It will use a seed that generates a number from 0 to 99.
// Sum the x + 1 with this random number.
// It returns the a integer random number.
func GenerateRandomNumbers(x int) int {
	return x + seed.Intn(100) + 1
}
