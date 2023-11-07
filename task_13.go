package main

import "fmt"

func main() {
	first, second := 22, 11
	fmt.Printf("first - %d, second - %d\n", first, second)

	first, second = second, first
	fmt.Printf("first - %d, second - %d\n", first, second)
}
