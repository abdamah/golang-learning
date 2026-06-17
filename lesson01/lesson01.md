# Lesson 01 — Hello World

**Author:** Abdallah  
**Copyright:** © 2026 Abdallah

---

## What you learn

Your first Go program. You will write code, run it, and see text printed in the terminal (the black window where commands run).

By the end of this lesson you will understand the **minimum** every Go program needs to work.

---

## Definitions

| Term | Simple meaning |
|------|----------------|
| **Program** | A file (or files) of instructions the computer follows. |
| **Terminal** | Where you type commands like `go run`. Also called command line or console. |
| **Package** | A way to organize Go code. Think of it as a folder label for your code. |
| **Import** | Bring in code someone else already wrote (from Go's standard library). |
| **Function** | A named block of code that does one job. |
| **`main`** | The special function where your program **starts**. |
| **String** | Text inside double quotes, like `"Hello"`. |

---

## Why this matters

Every Go application — from a tiny script to a big web server — starts with `package main`, imports, and `func main()`. Once you know this pattern, every future lesson will feel familiar.

---

## Key concepts

| Concept | Meaning |
|---------|---------|
| `package main` | This file is a **runnable program** (not a library others import). |
| `import "fmt"` | Use the **format** package — it can print text to the screen. |
| `func main()` | Go runs this function **first** when you start the program. |
| `fmt.Println()` | Print a line of text, then move to the next line. |

---

## Code walkthrough

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Author: Abdallah")
}
```

**Line by line:**

1. `package main` — tells Go: "this is an executable program."
2. `import "fmt"` — load the printing tools. `fmt` = **f**ormat.
3. `func main()` — the starting point. Empty `()` means this function takes no input.
4. `{` and `}` — everything inside runs when the program starts.
5. `fmt.Println("Hello, World!")` — print that text, then go to a new line.
6. Second `Println` — prints another line.

**Note:** Go uses **tabs** for indentation inside functions. Your editor usually adds them when you press Tab.

---

## Try it yourself (step by step)

1. Open `main.go` in your editor.
2. Change `"Abdallah"` to your own name in the second `Println`.
3. Save the file.
4. Run: `go run ./lesson01`
5. Confirm you see your name in the output.

---

## Run the lesson

From the project root:

```bash
go run ./lesson01
```

Or from inside the folder:

```bash
cd lesson01
go run .
```

**Expected output:**

```
Hello, World!
Author: Abdallah
```

---

## Common mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Forgetting quotes | `Println(Hello)` won't compile | Use `"Hello"` with double quotes |
| Wrong package name | `package Main` fails | Must be lowercase `package main` |
| No `func main()` | Program builds but does nothing useful | Every runnable app needs `main` |
| Running wrong folder | `go run` can't find files | `cd` into the lesson folder or use `go run ./lesson01` |

---

## Files in this lesson

| File | Purpose |
|------|---------|
| `main.go` | Lesson example — read and run this first |
| `lesson01.md` | This explanation |
| `exercise01.md` | Practice tasks |
| `solution01.go` | Answers (try exercises before opening) |

---

## Exercises

1. Read this file, then run `main.go`.
2. Open [exercise01.md](exercise01.md) and try:
   - Print your name.
   - Print your city.
   - Print `Learning Go`.
3. Check your work:

```bash
go run solution01.go
```

---

## Next lesson

**Lesson 02 — Data types** — learn `bool`, `int`, `float64`, and `string`.
