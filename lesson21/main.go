/*
 * Lesson 21 — JSON API
 *
 * Author:    Abdallah
 * Copyright: © 2026 Abdallah
 *
 * Description:
 *   GET /todos and POST /todos with JSON. Test in Postman.
 */

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

type Todo struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

type API struct {
	mu    sync.Mutex
	next  int
	todos []Todo
}

func (a *API) list(w http.ResponseWriter, r *http.Request) {
	a.mu.Lock()
	defer a.mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(a.todos)
}

func (a *API) create(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Text string `json:"text"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	if input.Text == "" {
		http.Error(w, "text is required", http.StatusBadRequest)
		return
	}

	a.mu.Lock()
	defer a.mu.Unlock()
	a.next++
	todo := Todo{ID: a.next, Text: input.Text}
	a.todos = append(a.todos, todo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func main() {
	api := &API{todos: []Todo{}}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /todos", api.list)
	mux.HandleFunc("POST /todos", api.create)

	log.Println("Server: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
