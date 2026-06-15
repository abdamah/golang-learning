// Run: go run solution17.go
//go:build ignore

package main

import (
	"encoding/json"
	"fmt"
)

type Note struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
	Done  bool   `json:"done"`
}

func main() {
	note := Note{ID: 1, Title: "Day 1", Body: "Started Go tutorial", Done: true}
	data, _ := json.Marshal(note)
	fmt.Println("JSON:", string(data))

	raw := `{"id":2,"title":"Walk","body":"30 min","done":false}`
	var decoded Note
	_ = json.Unmarshal([]byte(raw), &decoded)
	fmt.Println("Title:", decoded.Title, "Done:", decoded.Done)
}
