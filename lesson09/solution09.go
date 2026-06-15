// Run: go run solution09.go
//go:build ignore

package main

import "fmt"

func main() {
	ch := make(chan int, 2)

	go func() {
		for _, n := range []int{10, 20, 30} {
			ch <- n
		}
		close(ch)
	}()

	for n := range ch {
		fmt.Println(n)
	}
}
