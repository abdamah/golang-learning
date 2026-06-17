# Lesson 18 — First HTTP server

**Author:** Abdallah  
**Copyright:** © 2026 Abdallah

---

## What you learn

Turn your Go program into a **web server** that listens on port 8080 and responds to browser or Postman requests.

Uses only the standard library `net/http` — no Fiber yet.

---

## Definitions

| Term | Simple meaning |
|------|----------------|
| **HTTP** | Protocol browsers use to talk to servers (GET, POST, etc.) |
| **Server** | Program that waits for requests and sends responses |
| **Handler** | Function that runs when a URL is requested |
| **Route** | URL path like `/`, `/about`, `/health` |
| **Port** | Door number on your computer — `8080` means `localhost:8080` |
| **localhost** | Your own machine — `127.0.0.1` |

---

## Why this matters

Every API you build (Sopify backend, Fiber, etc.) ultimately handles HTTP requests. This lesson shows the core idea before frameworks add convenience.

---

## Key concepts

```go
func home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome! Server by Abdallah")
}

http.HandleFunc("/", home)
http.ListenAndServe(":8080", nil)
```

| Piece | Role |
|-------|------|
| `http.ResponseWriter` (`w`) | Where you **write** the response body |
| `*http.Request` (`r`) | Info about the **incoming** request (method, URL, headers) |
| `HandleFunc` | Register: "when someone visits this path, call this function" |
| `ListenAndServe` | Start listening — **blocks** until error or shutdown |

Server keeps running until you press **Ctrl+C** in the terminal.

---

## Step by step (first time)

1. Run: `go run ./lesson18`
2. See log: `Server: http://localhost:8080`
3. Open browser: `http://localhost:8080/`
4. Or use Postman: GET `http://localhost:8080/health`
5. Press Ctrl+C to stop the server

---

## Test in Postman

| Method | URL | Expected |
|--------|-----|----------|
| GET | `http://localhost:8080/` | Welcome message |
| GET | `http://localhost:8080/about` | Learning Go HTTP |
| GET | `http://localhost:8080/health` | ok |

---

## Common mistakes

| Mistake | Fix |
|---------|-----|
| Port already in use | Stop other server on 8080 or change port |
| Forgot to run server before Postman | Start `go run` first |
| Wrong URL `localhost:8080` without `http://` | Use full URL in Postman |

---

## Files in this lesson

| File | Purpose |
|------|---------|
| `main.go` | Lesson example |
| `lesson18.md` | This explanation |
| `exercise18.md` | Practice tasks |
| `solution18.go` | Answers |

---

## Exercises

Open [exercise18.md](exercise18.md) — test routes in Postman, add `/contact`.

---

## Next lesson

**Lesson 19 — Query params and status codes**
