/*
 * Lesson 13 — Structs
 *
 * Author:    Abdallah
 * Copyright: © 2026 Abdallah
 *
 * Description:
 *   Define a struct to represent a user or note.
 */

package main

import "fmt"

type User struct {
	ID    int
	Name  string
	Email string
}

func main() {
	me := User{ID: 1, Name: "Abdallah", Email: "abdallah@example.com"}
	fmt.Printf("%+v\n", me)
}
