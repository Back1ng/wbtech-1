package main

import (
	"fmt"
	"strconv"
)

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

func setBit(num int64, i int, value bool) int64 {
	fmt.Println("Getting:", strconv.FormatInt(num, 2))
	if value {
		// оператор ИЛИ
		num |= 1 << i
	} else {
		// оператор И НЕ
		num &^= 1 << i
	}
	fmt.Println("Returning:", strconv.FormatInt(num, 2))
	return num
}

func main() {
	a := int64(10)
	a = setBit(a, 1, false)
	a = setBit(a, 0, true)
}
