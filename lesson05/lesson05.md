# Lesson 05 — For loops

**Author:** Abdallah  
**Copyright:** © 2026 Abdallah

---

## What you learn

Repeat code without copying it many times. Loop over lists (slices) to process each item.

---

## Definitions

| Term | Simple meaning |
|------|----------------|
| **Loop** | Repeat a block of code multiple times |
| **Slice** | A flexible list: `[]string{"Go", "HTTP"}` |
| **`range`** | Go keyword to loop over each item in a slice |
| **Index** | Position in a list, starting at `0` for the first item |
| **Iterate** | Visit each item one by one |

---

## Why this matters

You will loop over products, cart items, database rows, and API results. One loop replaces dozens of copy-pasted lines.

---

## Key concepts

**Count loop** — like `for` in C/Java, but Go uses only `for`:

```go
for i := 1; i <= 5; i++ {
    fmt.Println("Count:", i)
}
```

| Part | Meaning |
|------|---------|
| `i := 1` | start at 1 |
| `i <= 5` | keep going while true |
| `i++` | add 1 after each round |

**Range loop** — over a slice:

```go
subjects := []string{"Go", "HTTP", "JSON"}
for index, subject := range subjects {
    fmt.Printf("%d. %s\n", index+1, subject)
}
```

`index` is 0, 1, 2… We print `index+1` so humans see 1, 2, 3.

---

## Try it yourself

1. Change the loop to count 1 to 10.
2. Add `"Postman"` to the `subjects` slice.
3. Run and count how many lines print.

---

## Run the lesson

```bash
go run ./lesson05
```

---

## Common mistakes

| Mistake | Fix |
|---------|-----|
| Off-by-one errors | Remember index starts at `0` |
| Infinite loop `for i := 1; i <= 5; i--` | `i--` decreases — use `i++` to count up |
| Thinking Go has `while` | Use `for condition { }` instead — same thing |

---

## Files in this lesson

| File | Purpose |
|------|---------|
| `main.go` | Lesson example |
| `lesson05.md` | This explanation |
| `exercise05.md` | Practice tasks |
| `solution05.go` | Answers |

---

## Exercises

Open [exercise05.md](exercise05.md) — evens, friends list, sum 1–10.

---

## Next lesson

**Lesson 06 — Switch** — cleaner branching for many options.
