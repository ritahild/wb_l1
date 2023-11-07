package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	var syncedMap = sync.Map{}
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			syncedMap.Store(i, rand.Int())
			wg.Done()
		}(i)
	}
	wg.Wait()

	syncedMap.Range(func(k, v interface{}) bool {
		fmt.Println("key:", k, ", val:", v)
		return true
	})

}
