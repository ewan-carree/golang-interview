package main

import "fmt"

// Based on this snippet, what is the output of the following code?

func Hello() {
	fmt.Println("Hello there")
}

func main() {
	go Hello()
	Hello()
	go Hello()
}
