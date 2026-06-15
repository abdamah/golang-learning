/*
 * Lesson 11 — Functions
 *
 * Author:    Abdallah
 * Copyright: © 2026 Abdallah
 *
 * Description:
 *   Write functions that return values and errors.
 */

package main

import "fmt"

func greet(name string) string {
	return "Hello, " + name
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	fmt.Println(greet("Abdallah"))

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("10 / 2 =", result)
}
