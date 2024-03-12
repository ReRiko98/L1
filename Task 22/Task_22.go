package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Значения переменных a и b
	a := big.NewInt(2)
	b := big.NewInt(20)

	// Возводим a в степень b
	exp := new(big.Int).Exp(a, b, nil)
	fmt.Println("a^b:", exp)

	// Деление a на b
	quotient := new(big.Int).Div(exp, a)
	fmt.Println("a^b / a:", quotient)

	// Сложение a и b
	sum := new(big.Int).Add(a, b)
	fmt.Println("a + b:", sum)

	// Вычитание b из a
	diff := new(big.Int).Sub(a, b)
	fmt.Println("a - b:", diff)
}
