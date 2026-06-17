# Lesson 08 — Goroutines

**Author:** Abdallah  
**Copyright:** © 2026 Abdallah

---

## What you learn

Run multiple pieces of work **at the same time** using goroutines. Go makes concurrency easy with one keyword: `go`.

---

## Definitions

| Term | Simple meaning |
|------|----------------|
| **Concurrency** | Doing more than one thing at once (or appearing to) |
| **Goroutine** | A lightweight thread started by `go function()` |
| **Main goroutine** | The one that runs `func main()` |
| **Synchronization** | Making sure goroutines finish before program exits |

---

## Why this matters

When 100 users hit your API, the server handles many requests concurrently — often one goroutine per request. Lesson 22 starts the HTTP server in a goroutine so `main` can wait for shutdown.

---

## Key concepts

```go
go say("from goroutine 1")
go say("from goroutine 2")
go say("from goroutine 3")

time.Sleep(100 * time.Millisecond)
fmt.Println("main done")
```

- `go` runs `say` in the **background** — it does not wait for it to finish.
- `main` might reach `main done` before goroutines print — unless we wait.
- `time.Sleep` pauses `main` for 100ms so goroutines have time to run.

**Output order can change** each run — that is normal with concurrency.

---

## Real-world analogy

Think of `go` as asking three people to shout a message while you count to one. You might hear "main done" before all shouts finish if you do not wait.

---

## Try it yourself

1. Remove `time.Sleep` — run several times. Do goroutines always print?
2. Increase sleep to `1 * time.Second` — all three should print before "main done".
3. Add a fourth goroutine.

---

## Run the lesson

```bash
go run ./lesson08
```

---

## Common mistakes

| Mistake | Fix |
|---------|-----|
| Program exits before goroutines run | Use `time.Sleep`, channels, or `WaitGroup` (Lesson 10) |
| Loop variable bug `go func() { fmt.Println(i) }()` | Pass `i` as parameter: `go func(n int) { ... }(i)` |

---

## Files in this lesson

| File | Purpose |
|------|---------|
| `main.go` | Lesson example |
| `lesson08.md` | This explanation |
| `exercise08.md` | Practice tasks |
| `solution08.go` | Answers |

---

## Exercises

Open [exercise08.md](exercise08.md) — launch 3 workers, understand why sleep is needed.

---

## Next lesson

**Lesson 09 — Channels** — safe communication between goroutines.
