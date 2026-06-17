# Lesson 02 — Data types

**Author:** Abdallah  
**Copyright:** © 2026 Abdallah

---

## What you learn

Every value in Go has a **type**. Types tell the computer what kind of data you are storing (text, whole number, decimal, true/false).

You will declare variables with an explicit type using `var`.

---

## Definitions

| Term | Simple meaning |
|------|----------------|
| **Variable** | A named box that holds a value. Example: `age` holds `20`. |
| **Type** | What kind of value fits in the box: number, text, etc. |
| **bool** | Boolean — only `true` or `false`. |
| **int** | Integer — whole numbers like `0`, `20`, `-5` (no decimals). |
| **float64** | Decimal numbers like `1.75`, `3.14`. |
| **string** | Text in double quotes: `"Abdallah"`. |
| **Declare** | Create a variable and give it a name and type. |

---

## Why this matters

Go is a **typed** language. You cannot put `"hello"` into an `int` variable — the compiler stops you before the program runs. This prevents many bugs. APIs and databases also use types (price = number, name = text).

---

## Key concepts

| Type | Example | Use for |
|------|---------|---------|
| `bool` | `true`, `false` | yes/no, on/off, logged in or not |
| `int` | `20`, `-5`, `0` | age, count, stock quantity |
| `float64` | `1.75`, `9.99` | price, height, ratings |
| `string` | `"Abdallah"` | names, emails, messages |

**Declaration syntax:**

```go
var name type = value
var age int = 20
```

---

## Code walkthrough

```go
var isStudent bool = true
var age int = 20
var height float64 = 1.75
var name string = "Abdallah"
```

- `isStudent` is `bool` — only true or false.
- `age` is `int` — no decimal point.
- `height` is `float64` — allows decimals. `64` means 64-bit precision (standard for decimals in Go).
- `name` is `string` — text always uses double quotes `"`.

Then we print each with a label so output is easy to read.

---

## Try it yourself

1. Add `var price float64 = 19.99`
2. Print it: `fmt.Println("Price:", price)`
3. Run and confirm the output.

---

## Run the lesson

```bash
go run ./lesson02
```

**Expected output:**

```
Student: true
Age: 20
Height: 1.75
Name: Abdallah
```

---

## Common mistakes

| Mistake | Fix |
|---------|-----|
| `var age int = "20"` | Age must be a number without quotes: `20` |
| `var name string = Abdallah` | Strings need quotes: `"Abdallah"` |
| Using `float` instead of `float64` | In Go, use `float64` for decimals |

---

## Files in this lesson

| File | Purpose |
|------|---------|
| `main.go` | Lesson example |
| `lesson02.md` | This explanation |
| `exercise02.md` | Practice tasks |
| `solution02.go` | Answers |

---

## Exercises

Open [exercise02.md](exercise02.md) — add `language` and `years` variables.

---

## Next lesson

**Lesson 03 — Variables** — shorter ways to declare variables with `:=`.
