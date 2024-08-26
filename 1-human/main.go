package main

import "fmt"

// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

type Human struct {
	Name string
	Age  int
}

func (h Human) SayName() {
	fmt.Printf("Hello. My name is: %s\n", h.Name)
}

type Action struct {
	Human
}

func main() {
	human := Human{
		Name: "Bob",
		Age:  20,
	}

	action := Action{
		Human: human,
	}

	action.SayName()
}
