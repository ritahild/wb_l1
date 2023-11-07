package main

import (
	"fmt"
	"sync"
)

func square(num int) {
	fmt.Println(num * num)
}

func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	for _, num := range arr {
		wg.Add(1)
		go func(num int) {
			square(num)
			wg.Done()
		}(num)
	}
	wg.Wait()
}
