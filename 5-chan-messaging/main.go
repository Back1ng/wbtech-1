package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"
)

// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

func main() {
	ch := make(chan int)

	// Определяем, сколько по времени должна работать программа
	seconds := 4

	ctx := context.Background()
	// определяем таймаут для контекста, сколько программа должна отработать
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(seconds))
	defer cancel()

	// Запускаем горутинку, которая посчитает сколько работала программа
	go func() {
		start := time.Now()
		select {
		case <-ctx.Done():
			log.Printf("Program worked %f seconds.\n", time.Since(start).Seconds())
		}
	}()

	// Последовательно отправляем значения в канал
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- rand.Intn(100)
			}
		}
	}()

	// пока контекст не закрыт - читаем из канала
	go func() {
		for {
			select {
			case v := <-ch:
				fmt.Println(v)
			case <-ctx.Done():
				return
			}
		}
	}()

	<-ctx.Done()
	<-time.After(time.Millisecond * 10)
}
