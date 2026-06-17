# Lesson 17 — JSON

**Author:** Abdallah  
**Copyright:** © 2026 Abdallah

---

## What you learn

**JSON** (JavaScript Object Notation) is the standard format for web APIs. Go converts structs ↔ JSON with `encoding/json`.

---

## Definitions

| Term | Simple meaning |
|------|----------------|
| **JSON** | Text format: `{"id":1,"title":"Day 1"}` |
| **Marshal** | Go struct → JSON bytes (`json.Marshal`) |
| **Unmarshal** | JSON bytes → Go struct (`json.Unmarshal`) |
| **JSON tag** | `` `json:"title"` `` — name in JSON (can differ from field name) |
| **Encoder/Decoder** | Stream JSON for HTTP bodies (Lesson 21) |

---

## Why this matters

Postman sends JSON. Your Next.js frontend sends JSON. Stripe webhooks send JSON. Struct tags control the field names clients see.

---

## Key concepts

```go
type Note struct {
    ID    int    `json:"id"`
    Title string `json:"title"`
    Body  string `json:"body"`
}
```

**To JSON:**

```go
data, err := json.Marshal(note)
fmt.Println(string(data))
// {"id":1,"title":"Day 1","body":"Started Go tutorial"}
```

**From JSON:**

```go
var copy Note
json.Unmarshal(data, &copy)   // pass pointer so Unmarshal can fill fields
fmt.Println(copy.Title)
```

---

## Try it yourself

1. Add `Done bool` with `` `json:"done"` ``.
2. Marshal and print — see `true`/`false` in JSON.
3. Unmarshal a JSON string from exercise file.

---

## Run the lesson

```bash
go run ./lesson17
```

---

## Common mistakes

| Mistake | Fix |
|---------|-----|
| Unmarshal without pointer | Use `&copy` not `copy` |
| Lowercase field names | JSON only sees exported (capitalized) fields |
| Invalid JSON | Unmarshal returns error — check `err` |

---

## Files in this lesson

| File | Purpose |
|------|---------|
| `main.go` | Lesson example |
| `lesson17.md` | This explanation |
| `exercise17.md` | Practice tasks |
| `solution17.go` | Answers |

---

## Exercises

Open [exercise17.md](exercise17.md) — add `Done` field, unmarshal sample JSON.

---

## Next lesson

**Lesson 18 — First HTTP server** — your program becomes a web server.
