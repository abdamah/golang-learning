# Lesson 14 — Methods

**Author:** Abdallah  
**Copyright:** © 2026 Abdallah

---

## What you learn

A **method** is a function attached to a type. Use **pointer receivers** when the method must change the struct's data.

---

## Definitions

| Term | Simple meaning |
|------|----------------|
| **Method** | Function with a receiver: `func (c *Counter) Add()` |
| **Receiver** | The variable before the function name — the struct it belongs to |
| **Pointer receiver** | `*Counter` — can modify the original struct |
| **Value receiver** | `Counter` — works on a copy; read-only for big changes |

---

## Why this matters

`func (a *API) listTodos(...)` is a method on your API struct. Handlers often use methods to access database and shared state.

---

## Key concepts

```go
type Counter struct {
    value int
}

func (c *Counter) Add() {
    c.value++
}

func (c Counter) Value() int {
    return c.value
}
```

- `Add` uses `*Counter` — increments the real counter.
- `Value` uses `Counter` — only reads; copy is fine.

```go
c := &Counter{}   // pointer to Counter
c.Add()
c.Add()
fmt.Println(c.Value())  // 2
```

`&Counter{}` creates a pointer to a new Counter.

---

## Try it yourself

1. Add `Reset()` method that sets `value` to 0.
2. Call `Reset()` after two `Add()` calls — print 0.
3. Explain in a comment: why `Add` needs `*Counter`.

---

## Run the lesson

```bash
go run ./lesson14
```

**Expected output:**

```
Count: 2
```

---

## Common mistakes

| Mistake | Fix |
|---------|-----|
| Value receiver when mutating | Use pointer receiver `*T` |
| Calling method on nil pointer | Ensure struct is created: `&Counter{}` |

---

## Files in this lesson

| File | Purpose |
|------|---------|
| `main.go` | Lesson example |
| `lesson14.md` | This explanation |
| `exercise14.md` | Practice tasks |
| `solution14.go` | Answers |

---

## Exercises

Open [exercise14.md](exercise14.md) — `Reset()`, `BankAccount` with `Deposit`.

---

## Next lesson

**Lesson 15 — Slices and maps**
