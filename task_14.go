package main

import (
	"fmt"
)

func main() {
	types := []interface{}{"Code", 0.8, true, 1}
	for _, x := range types {
		fmt.Printf("%T\n", x)
	}
}
