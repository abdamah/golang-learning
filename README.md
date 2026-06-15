# Go HTTP Tutorial — From Zero to http.Server

**Author:** Abdallah  
**Copyright:** © 2026 Abdallah. All rights reserved.

## About

A hands-on Go tutorial for beginners. You start with `Hello, World!` and work through 22 small, self-contained lessons until you can build a real HTTP server with `net/http` — no third-party frameworks (no Gin, Fiber, or Chi).

Each lesson lives in its own folder with example code, exercises, and a solution. Lessons 01–07 cover basics and control flow. Lessons 08–10 cover concurrency (goroutines, channels, select). Lessons 11–17 cover functions, structs, JSON, and errors. Lessons 18–22 introduce HTTP servers, routing, a JSON API, and production-style setup with `http.Server` and graceful shutdown. HTTP lessons are tested with [Postman](https://www.postman.com/downloads/).

**Requirements:** Go 1.24+

Each lesson folder contains:

- `main.go` — lesson example
- `exerciseNN.md` — exercises
- `solutionNN.go` — solution (run separately)

## Setup

```bash
go mod init github.com/abdamah/golang-learning  # already done
go run ./lesson01
```

## Exercises

```bash
# 1. Read the exercises
cat lesson04/exercise04.md

# 2. Try on your own, then check the solution
go run ./lesson04/solution04.go
```

> Solution files use `//go:build ignore` so `go run ./lesson04` still runs only `main.go`.

## Lessons

| # | Topic | Run |
|---|--------|-----|
| 01 | Hello World | `go run ./lesson01` |
| 02 | Data types | `go run ./lesson02` |
| 03 | Variables | `go run ./lesson03` |
| 04 | If / else | `go run ./lesson04` |
| 05 | For loops | `go run ./lesson05` |
| 06 | Switch | `go run ./lesson06` |
| 07 | Infinite for, break, continue | `go run ./lesson07` |
| 08 | Goroutines | `go run ./lesson08` |
| 09 | Channels | `go run ./lesson09` |
| 10 | Select and WaitGroup | `go run ./lesson10` |
| 11 | Functions | `go run ./lesson11` |
| 12 | Strings | `go run ./lesson12` |
| 13 | Structs | `go run ./lesson13` |
| 14 | Methods | `go run ./lesson14` |
| 15 | Slices and maps | `go run ./lesson15` |
| 16 | Errors | `go run ./lesson16` |
| 17 | JSON | `go run ./lesson17` |
| 18 | First HTTP server | `go run ./lesson18` |
| 19 | Query params and status codes | `go run ./lesson19` |
| 20 | Simple router | `go run ./lesson20` |
| 21 | JSON API (todos) | `go run ./lesson21` |
| 22 | http.Server (final) | `go run ./lesson22` |

## Postman (lessons 18–22)

1. Download [Postman](https://www.postman.com/downloads/)
2. Create a collection: **Go HTTP Tutorial**
3. Start the server: `go run ./lesson18` (or 19–22)
4. Send requests to `http://localhost:8080/...`
5. For JSON POST (lesson 21): Body → raw → JSON, header `Content-Type: application/json`
