package main

import (
	"fmt"
	"sort"
)

// Функция для выполнения бинарного поиска в отсортированном срезе
func binarySearch(arr []int, target int) int {
	// Сначала отсортируем срез, так как бинарный поиск требует отсортированного входного массива
	sort.Ints(arr)

	// Индексы начала и конца среза
	left := 0
	right := len(arr) - 1

	// Бинарный поиск
	for left <= right {
		// Находим середину среза
		mid := (left + right) / 2

		// Если элемент в середине среза равен искомому элементу, возвращаем его индекс
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			// Если элемент в середине меньше искомого элемента, сдвигаем левую границу
			left = mid + 1
		} else {
			// Если элемент в середине больше искомого элемента, сдвигаем правую границу
			right = mid - 1
		}
	}

	// Если искомый элемент не найден, возвращаем -1
	return -1
}

func main() {
	// Пример отсортированного среза
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Искомый элемент
	target := 5

	// Выполняем бинарный поиск
	index := binarySearch(arr, target)

	// Выводим результат
	if index != -1 {
		fmt.Printf("Element %d found at index %d\n", target, index)
	} else {
		fmt.Printf("Element %d not found in the array\n", target)
	}
}
