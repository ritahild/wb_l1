package main

import (
	"fmt"
)

//var justString string
//func someFunc() {
//	v := createHugeString(1 << 10)
//	justString = v[:100]
//}
//
//func main() {
//	someFunc()
//}

func createHugeString() string {
	s := "n"
	for i := 0; i < 10; i++ {
		s += s
	}
	return s
}

func someFunc() string {
	return createHugeString()[:100]
}

func main() {

	fmt.Println(someFunc())

}
