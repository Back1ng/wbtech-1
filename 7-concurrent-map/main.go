package main

import (
	"fmt"
	"sync"
)

// Реализовать конкурентную запись данных в map.

// SafeMap потокобезопасная обертка над мапой
type SafeMap struct {
	m  map[int]int
	mu sync.RWMutex
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		m: make(map[int]int),
	}
}

func (m *SafeMap) Put(k int, v int) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.m[k] = v
}

func (m *SafeMap) Get(k int) (v int) {
	// Используем Read Lock, когда пытаемся читать из мапы
	// Выгодно использовать при большом объеме чтения из-за отсутствия небольшой дополнительной проверки внутри метода
	m.mu.RLock()
	defer m.mu.RUnlock()

	v, ok := m.m[k]
	if ok {
		return v
	}

	return -1
}

func (m *SafeMap) Increment(k int) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.m[k]++
}

func main() {
	m := NewSafeMap()

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			m.Increment(0)
		}()
	}

	wg.Wait()
	fmt.Println(m.Get(0))
}
