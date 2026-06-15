// Run: go run solution16.go
//go:build ignore

package main

import (
	"errors"
	"fmt"
)

func validateUsername(name string) error {
	if name == "" {
		return errors.New("username is empty")
	}
	if len(name) < 3 {
		return errors.New("username too short")
	}
	return nil
}

func main() {
	for _, name := range []string{"", "ab", "abdallah"} {
		if err := validateUsername(name); err != nil {
			fmt.Println("FAIL:", name, "→", err)
		} else {
			fmt.Println("OK:", name)
		}
	}
}
