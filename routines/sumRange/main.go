package main

import (
	"fmt"
	"runtime"
)

//Create a function sumRange(start int, end int, resultChan chan int) that calculates the sum of integers from start to end and sends the result to a channel.
//Divide the range from 1 to N into M parts (no load balance is necessary), where M is the number of available CPU cores.
//Start M goroutines to calculate the sum of each part concurrently.
//Collect the results from all goroutines and calculate the final sum.
//Print the final sum.

// Example: 
// Given N = 10 (range from 1 to 10) and 4 available CPU cores.
// The range is divided into 4 parts: [1, 2], [3, 4], [5, 6], [7, 10].
// Each part is processed concurrently by a separate goroutine.
// Goroutines calculate the sum of their assigned part and send the result to a channel.
// The main goroutine collects the results from the channels and calculates the final sum.
// The final sum is printed to the console.

func main() {
	M := runtime.NumCPU()
	fmt.Println("Number of CPUs : ", M)

	N := 16 // 16 expects 136 as result
	// N : 10 // 10 expects 55 as result
	fmt.Printf("Range 1 to %d \n", N)
}
