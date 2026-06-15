// Run: go run solution06.go
//go:build ignore

package main

import "fmt"

func main() {
	status := 404
	switch status {
	case 200:
		fmt.Println("OK")
	case 404:
		fmt.Println("Not found")
	default:
		fmt.Println("Error")
	}

	role := "admin"
	switch role {
	case "admin":
		fmt.Println("Welcome, admin")
	case "user":
		fmt.Println("Welcome, user")
	default:
		fmt.Println("Welcome, guest")
	}

	method := "POST"
	switch method {
	case "GET":
		fmt.Println("Reading data")
	case "POST":
		fmt.Println("Creating a note")
	case "DELETE":
		fmt.Println("Deleting data")
	default:
		fmt.Println("Unknown method")
	}
}
