# Lesson 16 — Errors

**Author:** Abdallah  
**Copyright:** © 2026 Abdallah

---

## What you learn

Go does not have try/catch. Functions return an **`error`** as a second value. You check it with `if err != nil`.

---

## Definitions

| Term | Simple meaning |
|------|----------------|
| **error** | Built-in interface — anything that went wrong |
| **`nil`** | No error — success |
| **`errors.New`** | Create a simple error message |
| **`fmt.Errorf`** | Formatted error, can wrap another error with `%w` |
| **Wrap** | Add context: `fmt.Errorf("invalid port: %w", err)` |

---

## Why this matters

Database fails, JSON is invalid, port is wrong — every layer returns `error`. Ignoring errors causes silent bugs and bad API responses.

---

## Key concepts

```go
func parsePort(s string) (int, error) {
    if s == "" {
        return 0, errors.New("port is empty")
    }
    n, err := strconv.Atoi(s)
    if err != nil {
        return 0, fmt.Errorf("not a number: %w", err)
    }
    return n, nil
}
```

**Loop with continue on error:**

```go
port, err := parsePort(input)
if err != nil {
    fmt.Println("FAIL:", input, "→", err)
    continue   // try next input
}
fmt.Println("OK:", input, "→", port)
```

---

## Try it yourself

1. Test with `"8080"`, `"abc"`, `""` — see OK vs FAIL.
2. Write `validateAge` that errors if age < 0.
3. Always handle `err` before using other return values.

---

## Run the lesson

```bash
go run ./lesson16
```

---

## Common mistakes

| Mistake | Fix |
|---------|-----|
| `_` to ignore errors | Only ignore if you truly do not care |
| `if err != nil { }` empty block | Log, return, or handle — do not swallow silently |
| Returning wrong zero value on error | Return `0, err` for int — not a valid port |

---

## Files in this lesson

| File | Purpose |
|------|---------|
| `main.go` | Lesson example |
| `lesson16.md` | This explanation |
| `exercise16.md` | Practice tasks |
| `solution16.go` | Answers |

---

## Exercises

Open [exercise16.md](exercise16.md) — `validateUsername` with tests.

---

## Next lesson

**Lesson 17 — JSON** — APIs speak JSON.
