/*
 * Lesson 08 — Goroutines
 *
 * Author:    Abdallah
 * Copyright: © 2026 Abdallah
 *
 * Description:
 *   Run functions concurrently with go. main must wait or it exits too early.
 */

package main

import (
	"fmt"
	"time"
)

func say(msg string) {
	fmt.Println(msg)
}

func main() {
	go say("from goroutine 1")
	go say("from goroutine 2")
	go say("from goroutine 3")

	time.Sleep(100 * time.Millisecond)
	fmt.Println("main done")
}
