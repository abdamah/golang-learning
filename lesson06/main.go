/*
 * Lesson 06 — Switch
 *
 * Author:    Abdallah
 * Copyright: © 2026 Abdallah
 *
 * Description:
 *   Choose between many options with switch.
 */

package main

import "fmt"

func main() {
	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("Start of the week")
	case "Friday":
		fmt.Println("Almost weekend")
	case "Saturday", "Sunday":
		fmt.Println("Weekend!")
	default:
		fmt.Println("Regular day")
	}

	method := "GET"
	switch method {
	case "GET":
		fmt.Println("Read data")
	case "POST":
		fmt.Println("Create data")
	case "DELETE":
		fmt.Println("Delete data")
	default:
		fmt.Println("Unknown method")
	}
}
