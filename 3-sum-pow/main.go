package main

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
)

// Дана последовательность чисел: 2,4,6,8,10.
// Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

// main Пример с использованием atomic
func main() {
	at := withAtomic()
	mutex := withMutex()
	channel := withChannel()

	fmt.Printf("Atomic: %d\n Mutex: %d\n Channel: %d\n", at, mutex, channel)
}

func withAtomic() int64 {
	// Создаем массив с предопределенными значениями
	arr := [...]int{2, 4, 6, 8, 10}

	// Инициализируем atomic, waitGroup
	var sum atomic.Int64
	var wg sync.WaitGroup

	// Проходим значения в массиве, добавляя в итератор waitGroup по единице
	// Также можно один раз добавить заранее, используя wg.Add(len(arr))
	for _, v := range arr {
		wg.Add(1)

		// Запускаем горутину с вычислениями
		go func(v int) {
			// По завершению выполнения горутины уменьшаем счетчик waitGroup
			defer wg.Done()
			pow := math.Pow(float64(v), 2)
			// Атомарно увеличиваем atomic int64
			sum.Add(int64(pow))
		}(v)
	}

	// Дожидаемся выполнения всех горутин
	wg.Wait()

	return sum.Load()
}

func withMutex() int64 {
	arr := [...]int{2, 4, 6, 8, 10}

	var counter int64
	// Используем обычный Mutex, так как чтения у нас почти нет, и он рекомендуется в большинстве случаев
	var mu sync.Mutex
	// Аналогично примеру с атомиком используем waitGroup
	var wg sync.WaitGroup
	for _, v := range arr {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			// блокируем мьютекс, чтобы другие горутины ждали его разблокировки
			mu.Lock()
			// по окончании вычислений открываем ресурс для других горутин
			defer mu.Unlock()
			pow := math.Pow(float64(v), 2)
			// потокобезопасно(за счет мьютекса) инкрементируем значение счетчика
			counter += int64(pow)
		}(v)
	}

	// дожидаемся выполнения горутин
	wg.Wait()

	return counter
}

func withChannel() int64 {
	arr := [...]int{2, 4, 6, 8, 10}

	var counter int64
	// инициируем канал с интами
	calcs := make(chan int64)

	go func() {
		// по окончании выполнения закрываем канал для выхода из for range
		defer close(calcs)
		for _, v := range arr {
			pow := math.Pow(float64(v), 2)
			// заполняем канал рассчитанными значениями
			calcs <- int64(pow)
		}
	}()

	// пока канал не закрыт, читаем из него
	for v := range calcs {
		counter += v
	}

	return counter
}
