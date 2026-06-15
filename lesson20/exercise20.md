# Lesson 20 — Exercises

**Author:** Abdallah

## Postman

| Method | URL | Body | Expected |
|--------|-----|------|----------|
| GET | http://localhost:8080/ | — | Home |
| POST | http://localhost:8080/echo | hello postman | hello postman |
| GET | http://localhost:8080/nope | — | 404 |
| GET | http://localhost:8080/routes | — | route list |

## Exercises

1. Test echo and 404 in Postman.
2. Add **GET** `/routes` listing available paths.

## Check your solution

```bash
go run solution20.go
```
