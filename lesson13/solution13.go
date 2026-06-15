// Run: go run solution13.go
//go:build ignore

package main

import "fmt"

type User struct {
	ID    int
	Name  string
	Email string
}

type Note struct {
	ID    int
	Title string
	Body  string
}

func describeUser(u User) string {
	return fmt.Sprintf("%s (%s)", u.Name, u.Email)
}

func main() {
	me := User{ID: 1, Name: "Abdallah", Email: "abdallah@example.com"}
	fmt.Println(describeUser(me))
	fmt.Printf("%+v\n", Note{ID: 1, Title: "Shopping", Body: "Buy milk"})
	fmt.Printf("%+v\n", Note{ID: 2, Title: "Study", Body: "Lesson 13"})
}
