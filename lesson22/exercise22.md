# Lesson 22 — Exercises

**Author:** Abdallah

## Postman

| Method | URL | Expected |
|--------|-----|----------|
| GET | http://localhost:8080/ | Server by Abdallah |
| GET | http://localhost:8080/health | ok |
| GET | http://localhost:8080/version | 1.0.0 |
| GET | http://localhost:8080/author | Abdallah |
| GET | http://localhost:8080/info | JSON info |

## Exercises

1. Test all routes in Postman; watch terminal logs.
2. Press **Ctrl+C** — confirm graceful shutdown.
3. Add **GET** `/info` returning `{"author":"Abdallah","lesson":22}`.

## Check your solution

```bash
go run solution22.go
```
