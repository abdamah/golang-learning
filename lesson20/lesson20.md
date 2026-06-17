# Lesson 20 — Simple router

**Author:** Abdallah  
**Copyright:** © 2026 Abdallah

---

## What you learn

Build a small **router** that sends each request to the right handler based on **HTTP method + path**. Frameworks like Fiber do this for you — here you see how it works inside.

---

## Definitions

| Term | Simple meaning |
|------|----------------|
| **Router** | Matches request to handler: `GET /about` → about function |
| **Handler** | Function that handles one route |
| **`http.Handler`** | Interface with one method: `ServeHTTP(w, r)` |
| **404 Not Found** | No route matched this method + path |
| **Request body** | Data sent with POST — read with `io.ReadAll(r.Body)` |

---

## Why this matters

Sopify will have `GET /products`, `POST /cart`, etc. Understanding routing helps you debug "why did I get 404?" and design clean APIs.

---

## Key concepts

```go
type router struct {
    routes map[string]http.HandlerFunc
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
```

- Key is `"GET /"`, `"POST /echo"`, etc.
- `ok` from map lookup means route exists.
- Pass `rt` to `ListenAndServe` — it implements `http.Handler`.

**Echo POST** — returns body back:

```go
body, _ := io.ReadAll(r.Body)
w.Write(body)
```

---

## Test in Postman

| Method | URL | Body | Expected |
|--------|-----|------|----------|
| GET | `http://localhost:8080/` | — | Home |
| GET | `http://localhost:8080/about` | — | About Abdallah |
| POST | `http://localhost:8080/echo` | `hello postman` | same text back |
| GET | `http://localhost:8080/nope` | — | **404** |

For POST in Postman: Body → raw → type text.

---

## Try it yourself

1. Add `GET /health` → `ok`.
2. POST different text to `/echo` each time.
3. Trace in code: what key is built for `POST /echo`?

---

## Run the lesson

```bash
go run ./lesson20
```

---

## Common mistakes

| Mistake | Fix |
|---------|-----|
| Trailing slash mismatch | `/about` ≠ `/about/` — be consistent |
| Reading body twice | Body is a stream — read once, save bytes if needed |

---

## Files in this lesson

| File | Purpose |
|------|---------|
| `main.go` | Lesson example |
| `lesson20.md` | This explanation |
| `exercise20.md` | Practice tasks |
| `solution20.go` | Answers |

---

## Exercises

Open [exercise20.md](exercise20.md) — echo test, add `/routes` listing.

---

## Next lesson

**Lesson 21 — JSON API** — REST-style todos.
