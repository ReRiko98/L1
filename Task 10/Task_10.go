package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	// Данная последовательность температурных колебаний
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Создаем карту для хранения групп температур
	groups := make(map[int][]float64)

	// Заполняем карту данными
	for _, temp := range temperatures {
		// Определяем ключ (диапазон температур) для текущей температуры
		key := int(math.Floor(temp/10) * 10)

		// Добавляем текущую температуру в соответствующую группу
		groups[key] = append(groups[key], temp)
	}

	// Сортируем ключи карты для вывода в порядке возрастания
	var keys []int
	for key := range groups {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	// Выводим группы температурных значений
	for _, key := range keys {
		fmt.Printf("%d: %v\n", key, groups[key])
	}
}
