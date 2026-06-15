/*
 * Lesson 18 — First HTTP server
 *
 * Author:    Abdallah
 * Copyright: © 2026 Abdallah
 *
 * Description:
 *   Listen on :8080 and respond to GET / with plain text.
 *   Test in Postman: GET http://localhost:8080/
 */

package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome! Server by Abdallah")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Learning Go HTTP")
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ok")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/health", health)

	log.Println("Server: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
