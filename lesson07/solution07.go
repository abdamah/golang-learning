// Run: go run solution07.go
//go:build ignore

package main

import "fmt"

func main() {
	for i := 1; ; i++ {
		if i == 7 {
			fmt.Println("Done")
			break
		}
		fmt.Println(i)
	}

	fmt.Println("---")

	for i := 1; i <= 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

	// A web server runs in a loop to keep listening for new HTTP requests
	// until the program is stopped.
}
