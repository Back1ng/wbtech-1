package main

import (
	"fmt"
	"sort"
)

func main() {
	str := []int{1, 3, 5, 7, 9}

	searchable := 7

	res := sort.Search(len(str), func(i int) bool {
		return str[i] >= searchable
	})

	fmt.Println(res)
}
