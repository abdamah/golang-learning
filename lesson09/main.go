/*
 * Lesson 09 — Channels
 *
 * Author:    Abdallah
 * Copyright: © 2026 Abdallah
 *
 * Description:
 *   Send and receive values between goroutines using channels.
 */

package main

import "fmt"

func main() {
	// Unbuffered channel
	done := make(chan string)
	go func() {
		done <- "work finished"
	}()
	fmt.Println(<-done)

	// Buffered channel
	nums := make(chan int, 3)
	nums <- 1
	nums <- 2
	nums <- 3
	close(nums)

	fmt.Println("Buffered:")
	for n := range nums {
		fmt.Println(n)
	}
}
