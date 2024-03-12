package main

import (
	"fmt"
	"strings"
)

// Функция для переворачивания слов в строке
func reverseWords(input string) string {
	// Разбиваем строку на слова
	words := strings.Fields(input)

	// Инициализируем новый слайс для перевернутых слов
	reversedWords := make([]string, len(words))

	// Переворачиваем каждое слово и сохраняем его в новом слайсе
	for i, word := range words {
		reversedWords[len(words)-1-i] = word
	}

	// Объединяем перевернутые слова в одну строку
	reversedString := strings.Join(reversedWords, " ")

	return reversedString
}

func main() {
	// Пример строки
	str := "snow dog sun"

	// Выводим исходную строку
	fmt.Println("Исходная строка:", str)

	// Переворачиваем слова в строке
	reversed := reverseWords(str)

	// Выводим строку с перевернутыми словами
	fmt.Println("Перевернутая строка:", reversed)
}
