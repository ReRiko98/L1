package main

import (
	"fmt"
	"strings"
)

// Функция проверки уникальности символов в строке (регистронезависимо)
func areAllCharactersUnique(str string) bool {
	// Приводим строку к нижнему регистру
	str = strings.ToLower(str)

	// Создаем мапу для подсчета количества встреченных символов
	charCount := make(map[rune]int)

	// Итерируемся по каждому символу в строке
	for _, char := range str {
		// Если символ уже встречался, возвращаем false
		if charCount[char] > 0 {
			return false
		}
		// Увеличиваем счетчик для текущего символа
		charCount[char]++
	}

	// Если все символы уникальны, возвращаем true
	return true
}

func main() {
	// Примеры строк для проверки
	str1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "aabcd"

	// Проверяем уникальность символов в каждой строке и выводим результат
	fmt.Println("Уникальные символы в строке 'abcd':", areAllCharactersUnique(str1))
	fmt.Println("Уникальные символы в строке 'abCdefAaf':", areAllCharactersUnique(str2))
	fmt.Println("Уникальные символы в строке 'aabcd':", areAllCharactersUnique(str3))
}
