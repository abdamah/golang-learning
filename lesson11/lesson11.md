# Lesson 11 — Functions

**Author:** Abdallah  
**Copyright:** © 2026 Abdallah

---

## What you learn

Put reusable logic in **functions**. HTTP handlers are functions. Go functions can return **multiple values** — usually a result and an `error`.

---

## Definitions

| Term | Simple meaning |
|------|----------------|
| **Function** | Named block: `func name(params) returnType { }` |
| **Parameter** | Input to a function: `name string` |
| **Return value** | Output sent back to the caller |
| **error** | Built-in type for "something went wrong" |
| **`nil`** | Means "no error" or "empty" for pointers |

---

## Why this matters

You will write `func getProduct(id int) (Product, error)` and handlers like `func listProducts(w, r)`. The `if err != nil` pattern appears in almost every Go file.

---

## Key concepts

**Simple function:**

```go
func greet(name string) string {
    return "Hello, " + name
}
```

**Function with error:**

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}
```

**Calling and checking:**

```go
result, err := divide(10, 2)
if err != nil {
    fmt.Println("error:", err)
    return
}
fmt.Println("10 / 2 =", result)
```

Always check `err` before using `result`.

---

## Try it yourself

1. Call `divide(10, 0)` — see the error branch run.
2. Write `func double(n int) int` that returns `n * 2`.
3. Print `double(5)`.

---

## Run the lesson

```bash
go run ./lesson11
```

---

## Common mistakes

| Mistake | Fix |
|---------|-----|
| Ignoring `err` | Always check: `if err != nil` |
| Wrong return count | `(float64, error)` must return two values |
| `return` without handling error in main | Print or return error up the chain |

---

## Files in this lesson

| File | Purpose |
|------|---------|
| `main.go` | Lesson example |
| `lesson11.md` | This explanation |
| `exercise11.md` | Practice tasks |
| `solution11.go` | Answers |

---

## Exercises

Open [exercise11.md](exercise11.md) — `celsiusToFahrenheit`, `validateAge`.

---

## Next lesson

**Lesson 12 — Strings** — split, trim, and check text.
