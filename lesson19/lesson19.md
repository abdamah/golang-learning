# Lesson 19 — Query params and status codes

**Author:** Abdallah  
**Copyright:** © 2026 Abdallah

---

## What you learn

- Read **query parameters** from URLs: `/hello?name=Abdallah`
- Return correct **HTTP status codes** (200 OK, 400 Bad Request, 405 Method Not Allowed)

---

## Definitions

| Term | Simple meaning |
|------|----------------|
| **Query parameter** | Extra data after `?` in URL: `?name=Abdallah&age=20` |
| **Status code** | Number telling client if request succeeded: 200, 400, 404, 500 |
| **GET** | Read data — should not change server state |
| **400 Bad Request** | Client sent invalid data (missing required field) |
| **405 Method Not Allowed** | Wrong HTTP method (POST when only GET allowed) |

---

## Why this matters

Real APIs validate input: missing `name`, wrong method, invalid JSON. Status codes tell Postman and frontends what happened without reading the body only.

---

## Key concepts

**Read query param:**

```go
name := r.URL.Query().Get("name")
if name == "" {
    http.Error(w, "name is required", http.StatusBadRequest)
    return
}
```

**Check HTTP method:**

```go
if r.Method != http.MethodGet {
    http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
    return
}
```

`http.Error` sets status code **and** writes a message body.

---

## Test in Postman

| Method | URL | Expected |
|--------|-----|----------|
| GET | `http://localhost:8080/hello?name=Abdallah` | 200 + greeting |
| GET | `http://localhost:8080/hello` | **400** (no name) |
| GET | `http://localhost:8080/ping` | `pong` |
| GET | `http://localhost:8080/time` | current time |

In Postman, check the **Status** column (green 200 vs red 400).

---

## Try it yourself

1. Add `?name=` with empty value — should still be 400.
2. Try POST on `/hello` — add method check for 405.
3. Add `/ping` handler if not present.

---

## Run the lesson

```bash
go run ./lesson19
```

---

## Common mistakes

| Mistake | Fix |
|---------|-----|
| Forgetting `return` after `http.Error` | Handler keeps running and may write twice |
| Confusing path and query | `/hello/Abdallah` is path; `?name=Abdallah` is query |

---

## Files in this lesson

| File | Purpose |
|------|---------|
| `main.go` | Lesson example |
| `lesson19.md` | This explanation |
| `exercise19.md` | Practice tasks |
| `solution19.go` | Answers |

---

## Exercises

Open [exercise19.md](exercise19.md) — add `/bad` route, test in Postman.

---

## Next lesson

**Lesson 20 — Simple router**
