# Lesson 09 — Channels

**Author:** Abdallah  
**Copyright:** © 2026 Abdallah

---

## What you learn

Goroutines should not share memory carelessly. **Channels** let them send and receive values safely — like a pipe between goroutines.

---

## Definitions

| Term | Simple meaning |
|------|----------------|
| **Channel** | A typed pipe: `chan string`, `chan int` |
| **Send** | `ch <- value` — put value into channel |
| **Receive** | `value := <-ch` — take value from channel |
| **Unbuffered** | Sender waits until receiver is ready |
| **Buffered** | Channel holds N values before blocking |
| **Close** | `close(ch)` — no more sends; receivers can finish with `range` |

---

## Why this matters

Channels coordinate work: "when payment is done, send signal to create order". Lesson 22 uses a channel for OS shutdown signals (Ctrl+C).

---

## Key concepts

**Unbuffered — handoff:**

```go
done := make(chan string)
go func() {
    done <- "work finished"   // send
}()
fmt.Println(<-done)            // receive — waits until send happens
```

**Buffered — small queue:**

```go
nums := make(chan int, 3)   // buffer size 3
nums <- 1
nums <- 2
nums <- 3
close(nums)

for n := range nums {       // read until channel closed
    fmt.Println(n)
}
```

---

## Try it yourself

1. Change message to your name in the goroutine.
2. Try buffer size `1` and send 3 numbers — what blocks?
3. Print length of received values with a counter in `range`.

---

## Run the lesson

```bash
go run ./lesson09
```

**Expected output:**

```
work finished
Buffered:
1
2
3
```

---

## Common mistakes

| Mistake | Fix |
|---------|-----|
| Send on closed channel | Panic — only sender should close |
| Deadlock with unbuffered channel | Need both sender and receiver ready |
| Forgetting `close` before `range` | `range` on channel needs close to end |

---

## Files in this lesson

| File | Purpose |
|------|---------|
| `main.go` | Lesson example |
| `lesson09.md` | This explanation |
| `exercise09.md` | Practice tasks |
| `solution09.go` | Answers |

---

## Exercises

Open [exercise09.md](exercise09.md) — send numbers, buffered channel, close and range.

---

## Next lesson

**Lesson 10 — Select and WaitGroup**
