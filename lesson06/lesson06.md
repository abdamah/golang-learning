# Lesson 06 — Switch

**Author:** Abdallah  
**Copyright:** © 2026 Abdallah

---

## What you learn

When you have **many** options (HTTP methods, days of week, status codes), `switch` is cleaner than many `if else if` lines.

---

## Definitions

| Term | Simple meaning |
|------|----------------|
| **`switch`** | Compare one value against many cases |
| **`case`** | One possible match |
| **`default`** | Runs if nothing matched — like a final `else` |
| **Fallthrough** | In Go, only one `case` runs, then `switch` stops (no fall-through by default) |

---

## Why this matters

HTTP routers check method: GET, POST, DELETE. Status handlers check 200, 404, 500. `switch` makes this readable.

---

## Key concepts

```go
switch day {
case "Monday":
    fmt.Println("Start of the week")
case "Friday":
    fmt.Println("Almost weekend")
case "Saturday", "Sunday":    // multiple values in one case
    fmt.Println("Weekend!")
default:
    fmt.Println("Regular day")
}
```

**HTTP method example** — you will see this pattern in APIs:

```go
switch method {
case "GET":    // read
case "POST":   // create
case "DELETE": // remove
default:       // unknown
}
```

---

## Try it yourself

1. Change `day` to `"Sunday"` — run and see Weekend.
2. Change `method` to `"POST"` — see Create data.
3. Add `case "PUT":` with a print line.

---

## Run the lesson

```bash
go run ./lesson06
```

**Expected output:**

```
Start of the week
Read data
```

---

## Common mistakes

| Mistake | Fix |
|---------|-----|
| Forgetting `default` | Add `default` for unexpected values |
| Using `switch` for ranges | Use `if` for `score >= 80` or switch on true/false with if |

---

## Files in this lesson

| File | Purpose |
|------|---------|
| `main.go` | Lesson example |
| `lesson06.md` | This explanation |
| `exercise06.md` | Practice tasks |
| `solution06.go` | Answers |

---

## Exercises

Open [exercise06.md](exercise06.md) — status codes, roles, HTTP methods.

---

## Next lesson

**Lesson 07 — Infinite for, break, continue**
