# Lesson 19 — Exercises

**Author:** Abdallah

## Postman

| Method | URL | Expected |
|--------|-----|----------|
| GET | http://localhost:8080/hello?name=Abdallah | 200 |
| GET | http://localhost:8080/hello | 400 |
| GET | http://localhost:8080/ping | pong |
| GET | http://localhost:8080/bad | 400 |

## Exercises

1. Run all Postman tests above.
2. Add **GET** `/bad` that returns status 400.

## Check your solution

```bash
go run solution19.go
```
