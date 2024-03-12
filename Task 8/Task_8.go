package main

import (
	"fmt"
)

// Функция для установки i-го бита в указанное значение (1 или 0)
func setBit(num *int64, pos uint, bitValue uint) {
	// Создаем маску для установки бита
	mask := int64(1) << pos

	// Устанавливаем или сбрасываем бит в зависимости от значения bitValue
	if bitValue == 1 {
		*num |= mask // Установка бита в 1
	} else {
		*num &= ^mask // Сброс бита в 0
	}
}

func main() {
	// Исходное число
	var number int64 = 7 // Для примера

	// Устанавливаем 2-й бит в 1
	setBit(&number, 2, 1)
	fmt.Printf("Number after setting 2nd bit to 1: %d\n", number)

	// Устанавливаем 5-й бит в 0
	setBit(&number, 5, 0)
	fmt.Printf("Number after setting 5th bit to 0: %d\n", number)
}
