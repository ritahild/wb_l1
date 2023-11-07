package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func Worker(id int, jobs <-chan int) {
	for j := range jobs {
		time.Sleep(time.Second)
		fmt.Println("Worker", id, ": ", j)
	}
}

func GenerateJobs(jobs chan<- int, done <-chan bool) {
	for i := 1; true; i++ {
		select {
		case <-done:
			close(jobs)
			return
		default:
			fmt.Println(i)
			jobs <- rand.Intn(100)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func main() {

	const numWorkers = 5

	jobs := make(chan int, 100)

	wg := sync.WaitGroup{}

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			Worker(id, jobs)
		}(w)
	}

	done := make(chan bool, 1)

	go GenerateJobs(jobs, done)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	done <- true

	wg.Wait()

}
