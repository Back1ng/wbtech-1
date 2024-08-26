package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	counter atomic.Int64
}

func (c *Counter) Inc() {
	c.counter.Add(1)
}

func (c *Counter) Value() int64 {
	return c.counter.Load()
}

func main() {
	c := &Counter{}
	var wg sync.WaitGroup

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}

	wg.Wait()

	fmt.Println(c.Value())
}
