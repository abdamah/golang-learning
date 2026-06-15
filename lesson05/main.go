/*
 * Lesson 05 — For loops
 *
 * Author:    Abdallah
 * Copyright: © 2026 Abdallah
 *
 * Description:
 *   Count with for, loop over a slice with range.
 */

package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Count:", i)
	}

	subjects := []string{"Go", "HTTP", "JSON"}
	for index, subject := range subjects {
		fmt.Printf("%d. %s\n", index+1, subject)
	}
}
