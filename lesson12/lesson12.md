# Lesson 12 — Strings

**Author:** Abdallah  
**Copyright:** © 2026 Abdallah

---

## What you learn

Work with text using the `strings` package. URLs, paths, emails, and user input are all strings.

---

## Definitions

| Term | Simple meaning |
|------|----------------|
| **String** | Text in double quotes — immutable (cannot change in place) |
| **Trim** | Remove characters from start/end |
| **Split** | Cut string into parts by a separator |
| **Join** | Combine parts into one string |
| **Contains** | Check if substring exists inside string |

---

## Why this matters

API paths look like `/users/1/profile`. You split by `/` to get ID. You validate emails and search product names with string functions.

---

## Key concepts

```go
path := "/users/1/profile"
parts := strings.Split(strings.Trim(path, "/"), "/")
```

Step by step:

1. `strings.Trim(path, "/")` → `"users/1/profile"` (slashes removed from ends)
2. `strings.Split(..., "/")` → `["users", "1", "profile"]`
3. `strings.Contains(path, "users")` → `true`

---

## Try it yourself

1. Split `"go,http,server"` by comma.
2. Build `"/api/products"` with `strings.Join`.
3. Test `Contains` with a word that is not in the string — get `false`.

---

## Run the lesson

```bash
go run ./lesson12
```

---

## Common mistakes

| Mistake | Fix |
|---------|-----|
| Confusing `strings` package with string type | Import `"strings"` — lowercase package |
| Empty split parts | `"/a/b/"` may give empty strings — trim first |
| Case sensitivity | `"Go" != "go"` — use `strings.EqualFold` for case-insensitive |

---

## Files in this lesson

| File | Purpose |
|------|---------|
| `main.go` | Lesson example |
| `lesson12.md` | This explanation |
| `exercise12.md` | Practice tasks |
| `solution12.go` | Answers |

---

## Exercises

Open [exercise12.md](exercise12.md) — split, contains, join a path.

---

## Next lesson

**Lesson 13 — Structs** — group related data.
