/*
 * Lesson 19 — Query params and status codes
 *
 * Author:    Abdallah
 * Copyright: © 2026 Abdallah
 *
 * Description:
 *   Read query strings and return proper HTTP status codes.
 *   Test in Postman:
 *     GET http://localhost:8080/hello?name=Abdallah
 *     GET http://localhost:8080/hello          (expect 400)
 *     GET http://localhost:8080/time
 *     GET http://localhost:8080/ping
 */

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Hello, %s!\n", name)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, time.Now().Format("2006-01-02 15:04:05"))
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/ping", ping)

	log.Println("Server: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
