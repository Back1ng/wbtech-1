package main

import (
	"context"
	"log"
	"sync"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.

func main() {
	// Можно останавливать через канал отмены, контекст (с таймаутом, с отменой)
	// также используя флаг завершения

	var wg sync.WaitGroup
	wg.Add(3)
	cancelCh := make(chan struct{})
	go func() {
		defer wg.Done()
		select {
		case <-cancelCh:
			log.Println("Cancel channel closed.")
		}
	}()
	close(cancelCh)

	ctxWC, cancel := context.WithCancel(context.Background())
	go func() {
		defer wg.Done()
		select {
		case <-ctxWC.Done():
			log.Println("Context with cancel closed.")
		}
	}()
	cancel()

	ctxWT, cancel := context.WithTimeout(context.Background(), time.Second)
	go func() {
		defer wg.Done()
		select {
		case <-ctxWT.Done():
			log.Println("Context with timeout closed.")
		}
	}()

	<-ctxWT.Done()
	wg.Wait()
}
