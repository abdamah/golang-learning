/*
 * Lesson 15 — Slices and maps
 *
 * Author:    Abdallah
 * Copyright: © 2026 Abdallah
 *
 * Description:
 *   Store a todo list (slice) and status per item (map).
 */

package main

import "fmt"

func main() {
	todos := []string{"Learn Go", "Build server", "Test in Postman"}

	for i, todo := range todos {
		fmt.Printf("%d. %s\n", i+1, todo)
	}

	status := map[string]string{
		"Learn Go":     "done",
		"Build server": "in progress",
	}
	fmt.Println("First todo status:", status["Learn Go"])
}
