# Lesson 04 — If / else

**Author:** Abdallah  
**Copyright:** © 2026 Abdallah

---

## What you learn

Make your program **decide** what to do based on conditions. Real apps use this everywhere: "if password is empty, show error", "if user is admin, allow access".

---

## Definitions

| Term | Simple meaning |
|------|----------------|
| **Condition** | An expression that is true or false. Example: `age >= 18` |
| **Boolean expression** | Compares values: `>`, `<`, `>=`, `<=`, `==` (equal), `!=` (not equal) |
| **`if`** | Run code only when condition is true |
| **`else`** | Run different code when condition is false |
| **`else if`** | Check another condition if the first failed |

---

## Why this matters

Without `if/else`, programs do the same thing every time. Web servers use conditions to validate input, return errors, and check permissions.

---

## Key concepts

```go
if age >= 18 {
    fmt.Println("You can vote")
} else {
    fmt.Println("You are too young to vote")
}
```

**Rules in Go:**

- No parentheses around the condition — write `if age >= 18` not `if (age >= 18)`
- Braces `{ }` are **required** — even for one line
- Conditions must be `bool` — true or false

**Grade example** — checks from top to bottom, runs only **one** branch:

```go
if score >= 90       → Grade A
else if score >= 80  → Grade B
else                 → Grade C or below
```

---

## Try it yourself

1. Change `age` to `15` and run — see the else branch.
2. Change `score` to `95` — see Grade A.
3. Add: if `score < 0`, print `"Invalid score"`.

---

## Run the lesson

```bash
go run ./lesson04
```

**Expected output:**

```
You can vote
Grade: B
```

---

## Common mistakes

| Mistake | Fix |
|---------|-----|
| `if (age >= 18)` with parentheses | Remove parentheses — not needed in Go |
| `if age >= 18` without `{ }` | Always add braces on the next lines |
| Using `=` instead of `==` for compare | `=` assigns; `==` checks equality |

---

## Files in this lesson

| File | Purpose |
|------|---------|
| `main.go` | Lesson example |
| `lesson04.md` | This explanation |
| `exercise04.md` | Practice tasks |
| `solution04.go` | Answers |

---

## Exercises

Open [exercise04.md](exercise04.md) — hot day, password, and port checks.

---

## Next lesson

**Lesson 05 — For loops** — repeat code with `for` and `range`.
