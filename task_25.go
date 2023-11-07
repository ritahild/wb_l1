package main

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	timer := time.NewTimer(duration)
	<-timer.C
}

func main() {

	fmt.Println("Sleep")

	sleep(time.Second * 3)

	fmt.Println("Get")

}
