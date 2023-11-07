package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.Mutex
	counter int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.Lock()
	c.counter++
	c.Unlock()
}

func (c *Counter) String() string {
	return fmt.Sprint(c.counter)
}

func main() {

	c := NewCounter()
	var wg sync.WaitGroup

	for i := 0; i < 10000000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}

	wg.Wait()
	fmt.Println(c)
}
