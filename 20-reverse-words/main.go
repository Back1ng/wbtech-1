package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».
func main() {
	str := "snow dog sun"

	splitted := strings.Split(str, " ")
	for i, j := 0, len(splitted)-1; i < j; i, j = i+1, j-1 {
		splitted[i], splitted[j] = splitted[j], splitted[i]
	}

	fmt.Println(strings.Join(splitted, " "))
}
