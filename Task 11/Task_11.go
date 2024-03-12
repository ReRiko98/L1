package main

import "fmt"

// Функция для нахождения пересечения двух неупорядоченных множеств
func intersection(set1, set2 []int) []int {
	// Создаем карту для хранения элементов множества set1
	set1Map := make(map[int]bool)
	// Заполняем карту элементами из множества set1
	for _, num := range set1 {
		set1Map[num] = true
	}

	// Создаем слайс для хранения пересечения множеств
	intersect := make([]int, 0)

	// Проверяем каждый элемент множества set2 на наличие в множестве set1
	for _, num := range set2 {
		// Если элемент присутствует в обоих множествах, добавляем его в пересечение
		if set1Map[num] {
			intersect = append(intersect, num)
		}
	}

	return intersect
}

func main() {
	// Примеры неупорядоченных множеств
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{3, 4, 5, 6, 7}

	// Находим пересечение множеств
	result := intersection(set1, set2)

	// Выводим результат
	fmt.Println("Intersection:", result)
}
