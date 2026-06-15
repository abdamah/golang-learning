// Run: go run solution14.go
//go:build ignore

package main

import "fmt"

type Counter struct {
	value int
}

func (c *Counter) Add()  { c.value++ }
func (c *Counter) Reset() { c.value = 0 }
func (c Counter) Value() int { return c.value }

type BankAccount struct {
	balance float64
}

func (b *BankAccount) Deposit(amount float64) { b.balance += amount }
func (b BankAccount) Balance() float64        { return b.balance }

func main() {
	c := &Counter{}
	c.Add()
	c.Reset()
	fmt.Println("Count:", c.Value())

	account := &BankAccount{}
	account.Deposit(100)
	account.Deposit(50)
	fmt.Println("Balance:", account.Balance())
}
