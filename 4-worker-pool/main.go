package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров,
// которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.
// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

func main() {
	ch := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())

	// Реализуем Gracefully shutdown, нотифицируя канал о завершении
	done := make(chan os.Signal)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	workerCount := 2
	// Инициируем воркеров
	for i := 0; i < workerCount; i++ {
		go func(i int) {
			log.Printf("Worker #%d started.\n", i)

			for {
				select {
				case v := <-ch:
					fmt.Printf("Worker #%d get a value: %d\n", i, v)
				case <-ctx.Done():
					fmt.Printf("Worker #%d finished.\n", i)
					return
				}
			}
		}(i)
	}

	// Вижу несколько способ завершения воркеров
	// 1: Context
	// 2: Канал отмены
	// В целом, принцип их работы очень похож - мы закрываем канал, и получаем сигнал для завершения
	// Вариант с контекстом более гибкий

	// Создаем постоянный источник данных в канал
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

	// Как только мы получаем сигнал о завершении - закрываем контекст
	// который выключит воркеров и постоянный источник данных
	<-done

	cancel()
	<-time.After(time.Millisecond * 20)
}
