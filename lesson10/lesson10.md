# Lesson 10 — Select and WaitGroup

**Author:** Abdallah  
**Copyright:** © 2026 Abdallah

---

## What you learn

- **`select`** — wait on multiple channels; first ready wins
- **`sync.WaitGroup`** — wait until all goroutines finish

These two tools replace messy `Sleep` hacks from Lesson 08.

---

## Definitions

| Term | Simple meaning |
|------|----------------|
| **`select`** | Like `switch`, but for channels — waits for one case to be ready |
| **Timeout** | Give up after some time if nothing happens |
| **WaitGroup** | Counter: `Add` before work, `Done` when finished, `Wait` blocks until zero |
| **`defer`** | Run something when function exits — often `defer wg.Done()` |

---

## Why this matters

Lesson 22 shutdown: `select { case <-stop: }` waits for Ctrl+C. APIs use timeouts so slow clients do not hang forever.

---

## Key concepts

**Select with timeout:**

```go
select {
case msg := <-ready:
    fmt.Println(msg)
case <-time.After(2 * time.Second):
    fmt.Println("timeout")
}
```

If `ready` sends within 2 seconds → print message. Otherwise → print timeout.

**WaitGroup pattern:**

```go
var wg sync.WaitGroup
for i := 1; i <= 3; i++ {
    wg.Add(1)
    go func(id int) {
        defer wg.Done()
        fmt.Println("worker", id, "done")
    }(i)
}
wg.Wait()
fmt.Println("all workers finished")
```

1. `Add(1)` before each goroutine
2. `Done()` when goroutine finishes
3. `Wait()` in main until all are done

---

## Try it yourself

1. Change sleep to 3 seconds — see timeout.
2. Change loop to 5 workers.
3. Read Lesson 22 `select` on shutdown — same idea.

---

## Run the lesson

```bash
go run ./lesson10
```

---

## Common mistakes

| Mistake | Fix |
|---------|-----|
| `wg.Add` inside goroutine | Call `Add` before `go` |
| `wg.Add(0)` or forgetting `Done` | `Wait` blocks forever |
| Empty `select` with no `default` | Blocks forever — add timeout or case |

---

## Files in this lesson

| File | Purpose |
|------|---------|
| `main.go` | Lesson example |
| `lesson10.md` | This explanation |
| `exercise10.md` | Practice tasks |
| `solution10.go` | Answers |

---

## Exercises

Open [exercise10.md](exercise10.md) — timeout test, 5 workers, link to Lesson 22.

---

## Next lesson

**Lesson 11 — Functions** — reusable code and error returns.
