package main

import (
	"fmt"
)

// Функция для быстрой сортировки среза
func quicksort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)/2]
	left := make([]int, 0)
	right := make([]int, 0)
	equal := make([]int, 0)

	for _, value := range arr {
		switch {
		case value < pivot:
			left = append(left, value)
		case value == pivot:
			equal = append(equal, value)
		case value > pivot:
			right = append(right, value)
		}
	}

	left = quicksort(left)
	right = quicksort(right)

	return append(append(left, equal...), right...)
}

func main() {
	// Пример неотсортированного массива
	arr := []int{9, 4, 7, 2, 8, 5, 1, 6, 3}

	// Сортировка массива
	sorted := quicksort(arr)

	// Вывод отсортированного массива
	fmt.Println("Sorted array:", sorted)
}
