package main

import "fmt"

// Разработать программу, которая в рантайме способна определить
// тип переменной: int, string, bool, channel из переменной типа interface{}.
func main() {
	a := 123
	detect(a)
}

func detect(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("Given int")
	case string:
		fmt.Println("Given string")
	case bool:
		fmt.Println("Given bool")
	case chan int:
		fmt.Println("Given chan int")
	case chan string:
		fmt.Println("Given chan string")
	case chan bool:
		fmt.Println("Given chan bool")
	default:
		fmt.Printf("%#+v\n", v)
	}
}
