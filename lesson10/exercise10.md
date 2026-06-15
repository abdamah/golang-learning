# Lesson 10 — Exercises

**Author:** Abdallah

## Run the example

```bash
go run .
```

## Exercises

1. Change goroutine sleep to 3 seconds — confirm you get `timeout`.
2. Start 5 workers with `WaitGroup`, each printing its id.
3. How is `select` used in Lesson 22 for server shutdown?

## Check your solution

```bash
go run solution10.go
```

**Answer 3:** Lesson 22 uses `select` to wait for `SIGINT`/`SIGTERM` on a signal channel before shutting down.
