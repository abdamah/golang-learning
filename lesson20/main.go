/*
 * Lesson 20 — Simple router
 *
 * Author:    Abdallah
 * Copyright: © 2026 Abdallah
 *
 * Description:
 *   Route GET and POST requests to different handlers.
 *   Test in Postman:
 *     GET  http://localhost:8080/
 *     POST http://localhost:8080/echo  (body: hello postman)
 */

package main

import (
	"io"
	"log"
	"net/http"
)

type router struct {
	routes map[string]http.HandlerFunc
}

func newRouter() *router {
	return &router{routes: make(map[string]http.HandlerFunc)}
}

func (rt *router) Handle(method, path string, h http.HandlerFunc) {
	rt.routes[method+" "+path] = h
}

func (rt *router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + " " + r.URL.Path
	if h, ok := rt.routes[key]; ok {
		h(w, r)
		return
	}
	http.NotFound(w, r)
}

func main() {
	rt := newRouter()

	rt.Handle("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Home"))
	})
	rt.Handle("GET", "/about", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("About Abdallah"))
	})
	rt.Handle("GET", "/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	rt.Handle("POST", "/echo", func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		w.Write(body)
	})

	log.Println("Server: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", rt))
}
