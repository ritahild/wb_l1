package main

import (
	"context"
	"fmt"
	"time"
)

func Work() int {
	return 1
}

func main() {

	ch := make(chan int)

	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case ch <- Work():
				fmt.Println("Sending")
			case <-ctx.Done():
				close(ch)
				return
			}
			time.Sleep(time.Millisecond * 100)
		}
	}(ctx)

	go func() {
		time.Sleep(time.Second * 1)
		cancel()
	}()

	for data := range ch {
		fmt.Printf("Data %d\n", data)
	}

}
