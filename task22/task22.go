package task22

import (
	"fmt"
	"math/big"
)

const (
	firstString  = "4194304"
	secondString = "16777216"
)

// Тут просто идет работа с пакетом math/big
func RunTask22() {
	// записываем в переменные значения из строки, так как это проще
	first, _ := new(big.Int).SetString(firstString, 10)   // 2^22
	second, _ := new(big.Int).SetString(secondString, 10) // 2^24

	// оборачиваем методы math/big и используем их
	fmt.Println(multiply(first, second))
	fmt.Println(divide(second, first))
	fmt.Println(add(first, second))
	fmt.Println(subtract(second, first))
}

func multiply(a, b *big.Int) *big.Int {
	ans := new(big.Int)
	return ans.Mul(a, b)
}

func divide(a, b *big.Int) *big.Int {
	ans := new(big.Int)
	return ans.Div(a, b)
}

func add(a, b *big.Int) *big.Int {
	ans := new(big.Int)
	return ans.Add(a, b)
}

func subtract(a, b *big.Int) *big.Int {
	ans := new(big.Int)
	return ans.Sub(a, b)
}
