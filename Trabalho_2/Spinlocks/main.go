package main

import (
	syncprim "Trabalho_2/Spinlocks/syncprim"
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var acc int = 0
var thEnded = 0

//seed of random number
var seed = rand.New(rand.NewSource(time.Now().UnixNano()))

func generateRandomNumber() int8 {
	x := int8(seed.Intn(101))
	if seed.Float64() < 0.5 {
		return -x
	}
	return x
}

func sumThread(vec []int8) {
	var temp int
	for _, v := range vec {
		temp += int(v)
	}
	syncprim.Aquire()
	acc += temp
	thEnded++
	syncprim.Release()

}

func main() {
	var n int
	var k int
	flag.IntVar(&n, "n", 0, "Size of array N")
	flag.IntVar(&k, "k", 0, "Number of threads")
	flag.Parse()
	if n < 1 || k < 1 || n < k {
		fmt.Print("Incorrect flags values passed \n")
		return
	}
	//var compAcc = 0
	vector := make([]int8, n)
	for i := 0; i < n; i++ {
		x := generateRandomNumber()
		//compAcc += int(x)
		vector[i] = x
	}
	var avgTime float64
	for j := 0; j < 10; j++ {
		start := time.Now()
		i := 0
		for ; i < (n-n%k)-(n/k); i += (n / k) {
			go sumThread(vector[i:(i + n/k)])
		}
		go sumThread(vector[i:])
		for thEnded != k {
		}
		thEnded = 0
		duration := time.Since(start)
		avgTime += duration.Seconds()
	}
	fmt.Printf("Average Time Elapsed: %v seconds. For N:%v and k:%v \n", avgTime/10, n, k)
	fmt.Print("=====================\n")
}
