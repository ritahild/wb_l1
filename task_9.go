package main

import (
	"fmt"
	"sync"
)

func main() {
	num := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	ch1 := make(chan int)
	ch2 := make(chan int)

	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		for _, n := range num {
			ch1 <- n
		}
		close(ch1)
		wg.Done()
	}()

	go func() {
		for n := range ch1 {
			ch2 <- n * 2
		}
		close(ch2)
		wg.Done()
	}()

	go func() {
		for n := range ch2 {
			fmt.Println(n)
		}
		wg.Done()
	}()

	wg.Wait()

}
