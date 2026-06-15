/*
 * Lesson 14 — Methods
 *
 * Author:    Abdallah
 * Copyright: © 2026 Abdallah
 *
 * Description:
 *   Add methods to structs. Use pointers to update state.
 */

package main

import "fmt"

type Counter struct {
	value int
}

func (c *Counter) Add() {
	c.value++
}

func (c Counter) Value() int {
	return c.value
}

func main() {
	c := &Counter{}
	c.Add()
	c.Add()
	fmt.Println("Count:", c.Value())
}
