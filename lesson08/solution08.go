// Run: go run solution08.go
//go:build ignore

package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i <= 3; i++ {
		n := i
		go fmt.Println("worker", n)
	}

	time.Sleep(100 * time.Millisecond)
	fmt.Println("main done")

	// Without sleep, main can exit before goroutines run.
}
