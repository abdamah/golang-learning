// Run: go run solution12.go
//go:build ignore

package main

import (
	"fmt"
	"strings"
)

func main() {
	for _, topic := range strings.Split("go,http,server", ",") {
		fmt.Println(topic)
	}
	fmt.Println("Contains:", strings.Contains("hello world", "world"))
	fmt.Println("Path:", "/"+strings.Join([]string{"api", "notes"}, "/"))
}
