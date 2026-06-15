/*
 * Lesson 07 — Infinite loop, break, continue
 *
 * Author:    Abdallah
 * Copyright: © 2026 Abdallah
 *
 * Description:
 *   Use for{} with break and continue. Same idea as a server waiting for requests.
 */

package main

import "fmt"

func main() {
	count := 5
	for {
		fmt.Println("T-minus", count)
		count--
		if count == 0 {
			fmt.Println("Launch!")
			break
		}
	}

	fmt.Println("---")

	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println("Even:", i)
	}
}
