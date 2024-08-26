package main

import (
	"fmt"
	"math"
	"sync"
)

// Написать программу, которая конкурентно рассчитает значение
// квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

func main() {
	arr := [...]int{2, 4, 6, 8, 10}

	wg := sync.WaitGroup{}
	for _, v := range arr {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			fmt.Println(math.Pow(float64(v), 2))
		}(v)
	}

	wg.Wait()
}
