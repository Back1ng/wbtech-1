package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree)
// создать для нее собственное множество.
func main() {
	str := []string{"cat", "cat", "dog", "cat", "tree"}

	// Реализуем множество из уникальных значений
	set := make(map[string]bool)
	for _, v := range str {
		set[v] = true
	}

	// Создаем множество для последующего заполнения
	ownSet := make([]string, 0, len(set))

	// Заполняем множество уникальными значениями
	for k := range set {
		ownSet = append(ownSet, k)
	}

	fmt.Println(ownSet)
}
