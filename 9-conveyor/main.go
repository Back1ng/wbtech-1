package main

import (
	"fmt"
	"math"
	"sync"
)

// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
// во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
func main() {
	arr := [...]int{2, 4, 6, 8, 10}

	ch1, ch2 := make(chan int), make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		defer close(ch1)

		for i := 0; i < len(arr); i++ {
			ch1 <- arr[i]
		}
	}()

	go func() {
		defer wg.Done()
		defer close(ch2)
		for v := range ch1 {
			ch2 <- int(math.Pow(float64(v), 2))
		}
	}()

	for v := range ch2 {
		fmt.Println(v)
	}
}
