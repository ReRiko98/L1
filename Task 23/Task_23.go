package main

import "fmt"

// Функция для удаления i-ого элемента из слайса
func removeAtIndex(slice []int, index int) []int {
	// Проверяем, что индекс находится в допустимых пределах слайса
	if index < 0 || index >= len(slice) {
		return slice // Возвращаем оригинальный слайс без изменений
	}

	// Копируем элементы слайса до индекса и после индекса
	copy(slice[index:], slice[index+1:])

	// Уменьшаем длину слайса на 1
	return slice[:len(slice)-1]
}

func main() {
	// Исходный слайс
	slice := []int{1, 2, 3, 4, 5}

	// Удаляем элемент с индексом 2 (третий элемент)
	indexToRemove := 2
	result := removeAtIndex(slice, indexToRemove)

	// Выводим результат
	fmt.Println("Слайс после удаления элемента:", result)
}
