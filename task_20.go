package main

import (
	"fmt"
	"strings"
)

func reverse(str string) string {
	words := strings.Split(str, " ")
	var res []string

	for i := len(words) - 1; i >= 0; i-- {
		res = append(res, words[i])
	}

	return strings.Join(res, " ")
}

func main() {
	fmt.Println(reverse("snow dog sun"))
}
