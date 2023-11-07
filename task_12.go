package main

import "fmt"

func main() {
	animals := []string{"cat", "cat", "dog", "cat", "tree", "tree"}

	uniqueSlice := unique(animals)
	fmt.Println(uniqueSlice)
}

func unique(strSlice []string) []string {
	keys := make(map[string]bool)
	var unique []string
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			unique = append(unique, entry)
		}
	}
	return unique
}
