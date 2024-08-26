package main

import "fmt"

// Удалить i-ый элемент из слайса
func main() {
	var a = []int{1, 2, 3, 4, 5}
	fmt.Println(a)
	a = removeElement(a, 2)
	fmt.Println(a)
}

func removeElement(nums []int, key int) []int {
	if key < 0 || key >= len(nums) {
		return nums
	}

	return append(nums[:key], nums[key+1:]...)
}
