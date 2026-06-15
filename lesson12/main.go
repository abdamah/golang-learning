/*
 * Lesson 12 — Strings
 *
 * Author:    Abdallah
 * Copyright: © 2026 Abdallah
 *
 * Description:
 *   Trim, split, and check strings — useful for URLs and user input.
 */

package main

import (
	"fmt"
	"strings"
)

func main() {
	path := "/users/1/profile"
	parts := strings.Split(strings.Trim(path, "/"), "/")

	fmt.Println("Full path:", path)
	fmt.Println("Parts:", parts)
	fmt.Println("Contains 'users':", strings.Contains(path, "users"))
}
