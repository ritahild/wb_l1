package main

import (
	"fmt"
	"strconv"
)

func main() {
	var elems, lookingFor int

	fmt.Print("Количество ")
	fmt.Scan(&elems)

	var array = make([]int, elems)

	fmt.Print("Массив ")
	for i := 0; i < elems; i++ {
		array[i] = i
		fmt.Print(strconv.Itoa(i) + " ")
	}
	fmt.Println()

	fmt.Println("Поиск:")
	fmt.Scan(&lookingFor)

	var assumption, result = binarySearch(&array, lookingFor)
	if result {
		fmt.Println("Найдено: " + strconv.Itoa(assumption))
	} else {
		fmt.Println("Не найдено")
	}
}

func binarySearch(array *[]int, lookingFor int) (int, bool) {
	fmt.Println("Разыскивается")
	var mid, assumption int

	min := 0
	high := len(*array) - 1

	for min <= high {
		mid = (min + high) / 2
		assumption = mid
		if assumption == lookingFor {
			return assumption, true
		}
		if assumption > lookingFor {
			high = mid - 1
		} else {
			min = mid + 1
		}
	}
	return 0, false
}
