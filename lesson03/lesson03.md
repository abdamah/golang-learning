# Lesson 03 — Variables

**Author:** Abdallah  
**Copyright:** © 2026 Abdallah

---

## What you learn

Three ways to create variables in Go. In real programs you will use all three, but **`:=`** is the most common inside functions.

---

## Definitions

| Term | Simple meaning |
|------|----------------|
| **Declare** | Create a variable (give it a name). |
| **Assign** | Put a value into a variable. |
| **Infer** | Go figures out the type automatically from the value. |
| **`:=`** | Short declaration: declare + assign in one line (inside functions only). |
| **Zero value** | Default value when you declare without assigning: `0` for numbers, `""` for strings, `false` for bool. |

---

## Why this matters

Lesson 02 used the long form `var x int = 5`. That is clear for beginners but verbose. Go developers use `:=` daily because it is shorter and still safe — Go still knows the type.

---

## Key concepts

| Style | Example | When to use |
|-------|---------|-------------|
| Full `var` | `var city string = "Amman"` | You want to be explicit about the type |
| Inferred `var` | `var year = 2026` | Type is obvious; used at package level sometimes |
| Short `:=` | `country := "Jordan"` | **Most common** inside `func main()` and other functions |

**Important:** `:=` only works **inside functions**, not at the top level of a file (outside `main`).

---

## Code walkthrough

```go
var city string = "Amman"   // full form — type is written
country := "Jordan"         // short form — Go sees string
year := 2026                // short form — Go sees int
```

Go looks at `2026` and knows `year` is an `int`. It looks at `"Jordan"` and knows `country` is a `string`. You do not repeat the type.

---

## Try it yourself

1. Add `hobby := "reading"` inside `main`.
2. Add `var language string = "Go"`.
3. Print both. Run and check output.

---

## Run the lesson

```bash
go run ./lesson03
```

**Expected output:**

```
City: Amman
Country: Jordan
Year: 2026
```

---

## Common mistakes

| Mistake | Fix |
|---------|-----|
| `:=` outside a function | Use `var name = value` at package level instead |
| `city := "Amman"` then `city := "Irbid"` | Cannot declare same variable twice — use `city = "Irbid"` to reassign |
| Confusing `=` and `:=` | `:=` declares **new** variable; `=` updates existing one |

---

## Files in this lesson

| File | Purpose |
|------|---------|
| `main.go` | Lesson example |
| `lesson03.md` | This explanation |
| `exercise03.md` | Practice tasks |
| `solution03.go` | Answers |

---

## Exercises

Open [exercise03.md](exercise03.md) — practice `var` and `:=`.

---

## Next lesson

**Lesson 04 — If / else** — run code only when a condition is true.
