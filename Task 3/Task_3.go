package main

import (
	"fmt"
	"sync"
)

func main() {
	// Задаем массив чисел
	numbers := []int{2, 4, 6, 8, 10}

	// Создаем канал для хранения квадратов чисел
	squares := make(chan int)

	// Создаем WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Запускаем горутину для каждого числа
	for _, num := range numbers {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			// Возводим число в квадрат и отправляем результат в канал
			square := x * x
			squares <- square
		}(num)
	}

	// Запускаем горутину, которая закроет канал, когда все числа будут обработаны
	go func() {
		wg.Wait()
		close(squares)
	}()

	// Суммируем значения из канала
	sum := 0
	for square := range squares {
		sum += square
	}

	// Выводим сумму квадратов на стандартный вывод
	fmt.Println("Sum of squares:", sum)
}
