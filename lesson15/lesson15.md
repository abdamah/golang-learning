# Lesson 15 — Slices and maps

**Author:** Abdallah  
**Copyright:** © 2026 Abdallah

---

## What you learn

- **Slice** — ordered list you can grow with `append`
- **Map** — lookup table: key → value

Before PostgreSQL, many tutorials store data in slices and maps in memory.

---

## Definitions

| Term | Simple meaning |
|------|----------------|
| **Slice** | Dynamic list: `[]string{"a", "b"}` |
| **Map** | Dictionary / hash table: `map[string]string` |
| **Key** | What you look up: `"Learn Go"` |
| **Value** | What you get back: `"done"` |
| **append** | Add item to end of slice — returns new slice |

---

## Why this matters

Lesson 21 stores todos in a slice. Product lists are slices. Quick lookups (user ID → user) use maps.

---

## Key concepts

**Slice:**

```go
todos := []string{"Learn Go", "Build server", "Test in Postman"}
todos = append(todos, "Deploy app")
```

**Map:**

```go
status := map[string]string{
    "Learn Go":     "done",
    "Build server": "in progress",
}
fmt.Println(status["Learn Go"])  // done
```

If key does not exist, you get the **zero value** (`""` for strings) — not an error.

---

## Try it yourself

1. Loop map with `for key, value := range status`.
2. Add a new key to the map.
3. Print `len(todos)` after append.

---

## Run the lesson

```bash
go run ./lesson15
```

---

## Common mistakes

| Mistake | Fix |
|---------|-----|
| `append` without assignment | `todos = append(todos, item)` — append returns new slice |
| Expecting error on missing map key | Use `val, ok := m[key]` to check existence |

---

## Files in this lesson

| File | Purpose |
|------|---------|
| `main.go` | Lesson example |
| `lesson15.md` | This explanation |
| `exercise15.md` | Practice tasks |
| `solution15.go` | Answers |

---

## Exercises

Open [exercise15.md](exercise15.md) — append todo, add to map, print all entries.

---

## Next lesson

**Lesson 16 — Errors** — the `if err != nil` pattern.
