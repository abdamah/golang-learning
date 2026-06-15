/*
 * Lesson 17 — JSON
 *
 * Author:    Abdallah
 * Copyright: © 2026 Abdallah
 *
 * Description:
 *   Marshal and unmarshal JSON — same format Postman sends and receives.
 */

package main

import (
	"encoding/json"
	"fmt"
)

type Note struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func main() {
	note := Note{ID: 1, Title: "Day 1", Body: "Started Go tutorial"}

	data, err := json.Marshal(note)
	if err != nil {
		panic(err)
	}
	fmt.Println("JSON:", string(data))

	var copy Note
	if err := json.Unmarshal(data, &copy); err != nil {
		panic(err)
	}
	fmt.Println("Title:", copy.Title)
}
