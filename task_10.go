package main

import "fmt"

func main() {
	temperature := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groupedTemps := make(map[int][]float32)

	for _, temp := range temperature {
		key := roundToFormat(temp)
		groupedTemps[key] = append(groupedTemps[key], temp)
	}

	fmt.Println(groupedTemps)
}

func roundToFormat(val float32) int {

	return int(val) / 10 * 10
}
