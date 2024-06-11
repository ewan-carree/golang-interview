package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	content := strings.Repeat("A", 2000)
	os.WriteFile("testfile.txt", []byte(content), 0644)

	f, err := os.Open("testfile.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		if n == 0 {
			break
		}
		fmt.Print(string(buf[:n]))
	}
}
