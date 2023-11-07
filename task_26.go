package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(checkAllLettersUnique("abcd"))
	fmt.Println(checkAllLettersUnique("abCdefAaf"))
	fmt.Println(checkAllLettersUnique("aabcd"))
	fmt.Println(checkAllLettersUnique("aAbcd"))
}

func checkAllLettersUnique(s string) bool {
	s = strings.ToLower(s)
	var onceSeen []rune
	for _, symbol := range s {
		if !contains(onceSeen, symbol) {
			onceSeen = append(onceSeen, symbol)
		}
	}
	if len(onceSeen) != len(s) {
		return false
	}
	return true
}

func contains(s []rune, compareTo rune) bool {
	for _, a := range s {
		if a == compareTo {
			return true
		}
	}
	return false
}
