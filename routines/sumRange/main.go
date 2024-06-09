package main

import (
	"fmt"
	"runtime"
)

//Create a function sumRange(start int, end int, resultChan chan int) that calculates the sum of integers from start to end and sends the result to a channel.
//Divide the range from 1 to N into M parts, where M is the number of available CPU cores.
//Start M goroutines to calculate the sum of each part concurrently.
//Collect the results from all goroutines and calculate the final sum.
//Print the final sum.

func main() {
	M := runtime.NumCPU()
	fmt.Println("Number of CPUs : ", M)

	N := 16 // 16 expects 136 as result
	// N := 10 // 10 expects 55 as result
	fmt.Println("Range N : ", N)
}
