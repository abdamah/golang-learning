# Lesson 13 — Structs

**Author:** Abdallah  
**Copyright:** © 2026 Abdallah

---

## What you learn

A **struct** groups related fields into one type — like a row in a database or a JSON object.

---

## Definitions

| Term | Simple meaning |
|------|----------------|
| **Struct** | Custom type with named fields |
| **Field** | One piece of data inside a struct (`Name`, `Email`) |
| **Literal** | Create a value inline: `User{ID: 1, Name: "Abdallah"}` |
| **Struct tag** | Later: `` `json:"name"` `` tells JSON which key maps to field |

---

## Why this matters

`Product`, `User`, `Order` in Sopify are structs. JSON APIs convert structs ↔ JSON automatically.

---

## Key concepts

```go
type User struct {
    ID    int
    Name  string
    Email string
}

me := User{ID: 1, Name: "Abdallah", Email: "abdallah@example.com"}
fmt.Printf("%+v\n", me)
```

- `type User struct` defines a new type called `User`.
- Fields are capitalized if you want them **exported** (visible outside the package).
- `%+v` prints struct with field names: `{ID:1 Name:Abdallah ...}`

---

## Try it yourself

1. Add `Age int` to `User`.
2. Create a second user variable.
3. Print both with `%+v`.

---

## Run the lesson

```bash
go run ./lesson13
```

---

## Common mistakes

| Mistake | Fix |
|---------|-----|
| Lowercase field names for JSON API | Exported fields start with capital letter |
| Wrong field order in literal | Use `Name: value` syntax — order does not matter |

---

## Files in this lesson

| File | Purpose |
|------|---------|
| `main.go` | Lesson example |
| `lesson13.md` | This explanation |
| `exercise13.md` | Practice tasks |
| `solution13.go` | Answers |

---

## Exercises

Open [exercise13.md](exercise13.md) — `Note` struct, `describeUser` function.

---

## Next lesson

**Lesson 14 — Methods** — attach behavior to structs.
