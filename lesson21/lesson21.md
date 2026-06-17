# Lesson 21 — JSON API

**Author:** Abdallah  
**Copyright:** © 2026 Abdallah

---

## What you learn

Build a small **JSON API**: list todos (GET) and create todos (POST). Combine everything: structs, JSON, HTTP, methods, errors, and **mutex** for safe concurrent access.

---

## Definitions

| Term | Simple meaning |
|------|----------------|
| **REST API** | HTTP endpoints that create/read data, often using JSON |
| **GET** | Read data — return list of todos |
| **POST** | Create data — send new todo in body |
| **201 Created** | Status code meaning "resource was created" |
| **Mutex** | Lock so only one goroutine changes shared data at a time |
| **In-memory store** | Data lives in program memory — gone when server stops |

---

## Why this matters

This is a mini version of Sopify's product/cart API. Same patterns: decode JSON, validate, save, return JSON with correct status.

---

## Key concepts

**Routes (Go 1.22+ style):**

```go
mux.HandleFunc("GET /todos", api.list)
mux.HandleFunc("POST /todos", api.create)
```

**List — encode JSON:**

```go
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(a.todos)
```

**Create — decode JSON, validate:**

```go
var input struct {
    Text string `json:"text"`
}
json.NewDecoder(r.Body).Decode(&input)
if input.Text == "" {
    http.Error(w, "text is required", http.StatusBadRequest)
    return
}
```

**Mutex — thread safety:**

```go
a.mu.Lock()
defer a.mu.Unlock()
```

Two requests at once could corrupt the slice without a lock.

---

## Test in Postman

**1. List (empty first)**

- GET `http://localhost:8080/todos`
- Expected: `[]`

**2. Create todo**

- POST `http://localhost:8080/todos`
- Headers: `Content-Type: application/json`
- Body (raw JSON):

```json
{
  "text": "Test in Postman"
}
```

- Expected: status **201**, body `{"id":1,"text":"Test in Postman"}`

**3. List again** — should show your todo.

**4. Bad request** — POST `{"text":""}` → **400**

---

## Try it yourself

1. Create 3 todos — list shows array of 3.
2. POST invalid JSON — see 400.
3. Read `sync.Mutex` comment in code — why is it needed?

---

## Run the lesson

```bash
go run ./lesson21
```

---

## Common mistakes

| Mistake | Fix |
|---------|-----|
| Forgot `Content-Type: application/json` in Postman | Server may not parse body |
| No mutex on shared slice | Rare bugs under load — always lock |
| Returning 200 on create | Use `WriteHeader(201)` for POST create |

---

## Files in this lesson

| File | Purpose |
|------|---------|
| `main.go` | Lesson example |
| `lesson21.md` | This explanation |
| `exercise21.md` | Practice tasks |
| `solution21.go` | Answers |

---

## Exercises

Open [exercise21.md](exercise21.md) — create 3 todos, test 400 errors, add `/todos/count`.

---

## Next lesson

**Lesson 22 — http.Server** — production-style server (final lesson).
