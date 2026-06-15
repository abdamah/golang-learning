// Run: go run solution15.go
//go:build ignore

package main

import "fmt"

func main() {
	todos := []string{"Learn Go", "Build server", "Test in Postman"}
	todos = append(todos, "Deploy app")
	for i, todo := range todos {
		fmt.Printf("%d. %s\n", i+1, todo)
	}

	status := map[string]string{
		"Learn Go":        "done",
		"Build server":    "in progress",
		"Test in Postman": "pending",
	}
	for task, state := range status {
		fmt.Printf("%s: %s\n", task, state)
	}
}
