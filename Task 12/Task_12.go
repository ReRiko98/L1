package main

import "fmt"

// Функция для создания множества из строки
func createSet(strings []string) map[string]bool {
	set := make(map[string]bool)

	for _, str := range strings {
		set[str] = true
	}

	return set
}

func main() {
	// Исходная последовательность строк
	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создаем множество из строки
	set := createSet(strings)

	// Выводим содержимое множества
	fmt.Println("Set:", set)
}
