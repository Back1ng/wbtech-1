package main

import "fmt"

// Разработать программу, которая перемножает, делит, складывает, вычитает
// две числовых переменных a,b, значение которых > 2^20.
func main() {
	a := 1 << 21
	b := 1 << 23
	fmt.Println(Addition(a, b))
	fmt.Println(Subtraction(a, b))
	fmt.Println(Division(a, b))
	fmt.Println(Multiplication(a, b))
}

func Addition(a, b int) int {
	return a + b
}

func Subtraction(a, b int) int {
	return a - b
}

func Division(a, b int) int {
	return a / b
}

func Multiplication(a, b int) int {
	return a * b
}
