package main

import (
	"fmt"
)

func main() {
	// Создаем каналы для передачи данных
	inputChannel := make(chan int)
	outputChannel := make(chan int)

	// Горутина для записи чисел во второй канал с применением операции умножения на 2
	go func() {
		for x := range inputChannel {
			// Применяем операцию умножения на 2 к числу x и записываем результат во второй канал
			outputChannel <- x * 2
		}
		// Закрываем второй канал после завершения записи данных
		close(outputChannel)
	}()

	// Горутина для вывода данных из второго канала в stdout
	go func() {
		for result := range outputChannel {
			// Выводим результат операции в stdout
			fmt.Println("Result:", result)
		}
	}()

	// Данные из массива, которые будут отправлены в первый канал
	numbers := []int{1, 2, 3, 4, 5}

	// Записываем числа из массива в первый канал
	for _, num := range numbers {
		inputChannel <- num
	}

	// Закрываем первый канал после записи всех данных
	close(inputChannel)

	// Ожидаем завершения вывода данных в stdout
	select {}
}
