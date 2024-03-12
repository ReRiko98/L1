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
			square := x * x
			squares <- square // Отправляем квадрат числа в канал
		}(num)
	}

	// Запускаем горутину, которая закроет канал, когда все числа будут обработаны
	go func() {
		wg.Wait()
		close(squares)
	}()

	// Читаем значения из канала и выводим их на стандартный вывод
	for square := range squares {
		fmt.Println(square)
	}
}
