/*
 * Lesson 10 — Select and WaitGroup
 *
 * Author:    Abdallah
 * Copyright: © 2026 Abdallah
 *
 * Description:
 *   Wait on multiple channels with select. Wait for goroutines with sync.WaitGroup.
 */

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ready := make(chan string)
	go func() {
		time.Sleep(500 * time.Millisecond)
		ready <- "server ready"
	}()

	select {
	case msg := <-ready:
		fmt.Println(msg)
	case <-time.After(2 * time.Second):
		fmt.Println("timeout")
	}

	fmt.Println("---")

	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Println("worker", id, "done")
		}(i)
	}
	wg.Wait()
	fmt.Println("all workers finished")
}
