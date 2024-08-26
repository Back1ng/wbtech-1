package main

import (
	"fmt"
	"sort"
)

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка
func main() {
	a := []int{1, 3, 2, 4, 7, 1, 123, 6, 34, 31, 6, 3, 1, 15, 17, 14}
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	fmt.Println(a)
}
