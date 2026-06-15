/*
 * Lesson 04 — If / else
 *
 * Author:    Abdallah
 * Copyright: © 2026 Abdallah
 *
 * Description:
 *   Make decisions with if, else if, and else.
 */

package main

import "fmt"

func main() {
	age := 18

	if age >= 18 {
		fmt.Println("You can vote")
	} else {
		fmt.Println("You are too young to vote")
	}

	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else {
		fmt.Println("Grade: C or below")
	}
}
