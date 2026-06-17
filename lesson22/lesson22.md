# Lesson 22 — http.Server

**Author:** Abdallah  
**Copyright:** © 2026 Abdallah

---

## What you learn

The **final lesson**: run a real server with `http.Server` — timeouts, logging middleware, and **graceful shutdown** on Ctrl+C.

Brings together lessons **08–10** (goroutines, channels, select) with HTTP.

---

## Definitions

| Term | Simple meaning |
|------|----------------|
| **`http.Server`** | Struct with config: address, handler, timeouts |
| **Timeout** | Max time for read/write — protects against slow clients |
| **Middleware** | Function that wraps handlers (e.g. logging every request) |
| **Graceful shutdown** | Stop accepting new requests, finish active ones, then exit |
| **Signal** | OS message like Ctrl+C (`SIGINT`) |
| **Shutdown** | `srv.Shutdown(ctx)` — clean stop |

---

## Why this matters

`ListenAndServe` alone is fine for learning. Production apps use `http.Server` for timeouts and safe deploys (no cutting off active users mid-request).

---

## Key concepts

**Server with timeouts:**

```go
srv := &http.Server{
    Addr:         ":8080",
    Handler:      logging(mux),
    ReadTimeout:  5 * time.Second,
    WriteTimeout: 10 * time.Second,
    IdleTimeout:  60 * time.Second,
}
```

**Start in goroutine** (Lesson 08):

```go
go func() {
    srv.ListenAndServe()
}()
```

**Wait for Ctrl+C** (Lesson 09 + 10):

```go
stop := make(chan os.Signal, 1)
signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
<-stop   // block until signal

srv.Shutdown(ctx)
```

**Logging middleware** — runs before and after each handler:

```go
func logging(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
    })
}
```

---

## Step by step

1. Run: `go run ./lesson22`
2. Postman: GET `http://localhost:8080/health`
3. Terminal shows: `GET /health 0.5ms` (example)
4. Press **Ctrl+C**
5. See: `shutdown signal received` → `server stopped`

---

## Test in Postman

| Method | URL | Expected |
|--------|-----|----------|
| GET | `http://localhost:8080/` | Server by Abdallah |
| GET | `http://localhost:8080/health` | ok |
| GET | `http://localhost:8080/version` | 1.0.0 |
| GET | `http://localhost:8080/author` | Abdallah |

---

## Common mistakes

| Mistake | Fix |
|---------|-----|
| `ListenAndServe` in main without `go` | Blocks forever — cannot wait for signal |
| No timeout on Shutdown context | Use `context.WithTimeout` so shutdown does not hang |
| Forgetting `http.ErrServerClosed` | Normal after Shutdown — do not treat as fatal |

---

## Files in this lesson

| File | Purpose |
|------|---------|
| `main.go` | Lesson example |
| `lesson22.md` | This explanation |
| `exercise22.md` | Practice tasks |
| `solution22.go` | Answers |

---

## Exercises

Open [exercise22.md](exercise22.md) — test all routes, Ctrl+C shutdown, add `/info` JSON.

---

## What you completed

You went from **Hello World** to a production-style HTTP server using only the Go standard library.

**Next:** Sopify full-stack course — Fiber v3 backend, Next.js frontend, Clerk, Cloudflare Images, Stripe. See `sopify/` and `sopif_course_plan.md`.
