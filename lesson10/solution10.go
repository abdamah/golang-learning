// Run: go run solution10.go
//go:build ignore

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ready := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		ready <- "server ready"
	}()

	select {
	case msg := <-ready:
		fmt.Println(msg)
	case <-time.After(2 * time.Second):
		fmt.Println("timeout")
	}

	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Println("worker", id, "done")
		}(i)
	}
	wg.Wait()
	fmt.Println("all workers finished")
}
