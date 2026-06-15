// Run: go run solution05.go
//go:build ignore

package main

import "fmt"

func main() {
	fmt.Println("Even numbers:")
	for i := 2; i <= 10; i += 2 {
		fmt.Println(i)
	}

	friends := []string{"Ali", "Sara", "Omar"}
	fmt.Println("Friends:")
	for i, name := range friends {
		fmt.Printf("%d. %s\n", i+1, name)
	}

	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("Sum 1-10:", sum)
}
