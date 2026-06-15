// Run: go run solution04.go
//go:build ignore

package main

import "fmt"

func main() {
	temperature := 30
	if temperature > 25 {
		fmt.Println("Hot day")
	} else {
		fmt.Println("Nice day")
	}

	password := ""
	if password == "" {
		fmt.Println("Password required")
	} else {
		fmt.Println("Password OK")
	}

	port := 8080
	if port >= 1 && port <= 65535 {
		fmt.Println("Valid port")
	} else {
		fmt.Println("Invalid port")
	}
}
