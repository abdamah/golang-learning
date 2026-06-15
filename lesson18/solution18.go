// Run: go run solution18.go
//go:build ignore

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome! Server by Abdallah")
	})
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Learning Go HTTP")
	})
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})
	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Email: abdallah@example.com")
	})

	log.Println("Server: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
