package main

import "fmt"

// Реализовать пересечение двух неупорядоченных множеств
func main() {
	a := []int{1, 2, 3, 4}
	b := []int{3, 4, 4, 5, 6}
	c := make([]int, 0, len(a))

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				c = append(c, a[i])
			}
		}
	}

	fmt.Println(c)
}
