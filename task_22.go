package main

import (
	"fmt"
	"math/big"
)

func add(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

func sub(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

func multiply(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

func divide(a, b *big.Int) *big.Int {
	return new(big.Int).Div(a, b)
}

func main() {

	bigIntA := big.NewInt(int64(1 << 30))
	bigIntB := big.NewInt(int64(1 << 25))

	fmt.Println(bigIntA)
	fmt.Println(bigIntB)

	fmt.Println("Add: ", add(bigIntA, bigIntB))

	fmt.Println("Sub: ", sub(bigIntA, bigIntB))

	fmt.Println("Multiply: ", multiply(bigIntA, bigIntB))

	fmt.Println("Devide: ", divide(bigIntA, bigIntB))

}
