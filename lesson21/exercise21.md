# Lesson 21 — Exercises

**Author:** Abdallah

## Postman

| Method | URL | Body | Expected |
|--------|-----|------|----------|
| GET | http://localhost:8080/todos | — | `[]` |
| POST | http://localhost:8080/todos | `{"text":"Buy milk"}` | 201 |
| POST | http://localhost:8080/todos | `{"text":""}` | 400 |
| GET | http://localhost:8080/todos/count | — | `{"count":N}` |

## Exercises

1. Create 3 todos in Postman.
2. Test empty text and invalid JSON (400).
3. Add **GET** `/todos/count`.

## Check your solution

```bash
go run solution21.go
```
