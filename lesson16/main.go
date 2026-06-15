/*
 * Lesson 16 — Errors
 *
 * Author:    Abdallah
 * Copyright: © 2026 Abdallah
 *
 * Description:
 *   Return and check errors with if err != nil.
 */

package main

import (
	"errors"
	"fmt"
	"strconv"
)

func parsePort(s string) (int, error) {
	if s == "" {
		return 0, errors.New("port is empty")
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("not a number: %w", err)
	}
	return n, nil
}

func main() {
	for _, input := range []string{"8080", "abc", ""} {
		port, err := parsePort(input)
		if err != nil {
			fmt.Println("FAIL:", input, "→", err)
			continue
		}
		fmt.Println("OK:", input, "→", port)
	}
}
