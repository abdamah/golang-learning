# Lesson 07 — Infinite loop, break, continue

**Author:** Abdallah  
**Copyright:** © 2026 Abdallah

---

## What you learn

- Loop forever with `for { }` until you `break`
- Skip one round with `continue`
- Understand why servers "loop" waiting for requests

---

## Definitions

| Term | Simple meaning |
|------|----------------|
| **Infinite loop** | `for { }` with no condition — runs until `break` |
| **`break`** | Exit the loop immediately |
| **`continue`** | Skip rest of this round, go to next iteration |
| **Server loop** | A server keeps running, waiting for the next HTTP request |

---

## Why this matters

Web servers do not exit after one request. They loop: wait → handle request → wait again. Shutdown (Lesson 22) sends a signal to `break` that loop cleanly.

---

## Key concepts

**Countdown with break:**

```go
for {
    fmt.Println("T-minus", count)
    count--
    if count == 0 {
        fmt.Println("Launch!")
        break    // leave the loop
    }
}
```

**Skip odds with continue:**

```go
for i := 1; i <= 10; i++ {
    if i%2 != 0 {
        continue   // skip to next i — don't print odds
    }
    fmt.Println("Even:", i)
}
```

`i % 2` is the remainder when dividing by 2. Odd numbers have remainder 1.

---

## Try it yourself

1. Remove `break` temporarily — see infinite output (press Ctrl+C to stop).
2. Change `continue` to skip number 3 instead of all odds.
3. Write in a comment: "A server loops until shutdown."

---

## Run the lesson

```bash
go run ./lesson07
```

---

## Common mistakes

| Mistake | Fix |
|---------|-----|
| Infinite loop with no break | Always have an exit: `break`, `return`, or condition |
| `break` only breaks inner loop | Use labels or refactor if you need to break outer loop |

---

## Files in this lesson

| File | Purpose |
|------|---------|
| `main.go` | Lesson example |
| `lesson07.md` | This explanation |
| `exercise07.md` | Practice tasks |
| `solution07.go` | Answers |

---

## Exercises

Open [exercise07.md](exercise07.md) — break at 7, skip 5, explain servers and loops.

---

## Next lesson

**Lesson 08 — Goroutines** — run code concurrently.
