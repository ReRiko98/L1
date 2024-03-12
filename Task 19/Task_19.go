package main

import "fmt"

// Функция для переворачивания строки
func reverseString(input string) string {
	// Преобразуем строку в слайс рун
	runes := []rune(input)

	// Получаем длину строки
	length := len(runes)

	// Итерируемся до середины строки и меняем местами символы
	for i := 0; i < length/2; i++ {
		runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
	}

	// Возвращаем перевернутую строку
	return string(runes)
}

func main() {
	// Пример строки
	str := "главрыба — абырвалг"

	// Выводим исходную строку
	fmt.Println("Исходная строка:", str)

	// Переворачиваем строку
	reversed := reverseString(str)

	// Выводим перевернутую строку
	fmt.Println("Перевернутая строка:", reversed)
}
