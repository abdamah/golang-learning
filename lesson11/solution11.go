// Run: go run solution11.go
//go:build ignore

package main

import "fmt"

func celsiusToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

func validateAge(age int) error {
	if age < 0 || age > 120 {
		return fmt.Errorf("invalid age: %d", age)
	}
	return nil
}

func main() {
	fmt.Println("25°C =", celsiusToFahrenheit(25), "°F")

	if err := validateAge(20); err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("age 20 is valid")
	}
}
