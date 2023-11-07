package main

import "fmt"

func delete_at_index(slice []int, index int) []int {

	return append(slice[:index], slice[index+1:]...)
}

func main() {

	numbers := []int{10, 20, 30, 40, 50, 90, 60}
	fmt.Println("Original Slice:", numbers)

	var index int = 6

	elem := numbers[index]

	numbers = delete_at_index(numbers, index)

	fmt.Printf("The element %d was deleted.\n", elem)
	fmt.Println("Slice after deleting element:", numbers)
}
