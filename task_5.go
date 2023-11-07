package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	chanel := make(chan int)
	defer close(chanel)

	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			chanel <- rand.Intn(500)
		}
	}()

	go func() {
		for data := range chanel {
			fmt.Printf("recieved number %d\n", data)
		}
	}()

	time.Sleep(10 * time.Second)

}
