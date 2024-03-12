package main

import (
	"fmt"
	"time"
)

func main() {
	// Создаем канал для передачи данных
	dataChannel := make(chan int)

	// Время выполнения программы (в секундах)
	N := 5

	// Создаем горутину для отправки данных в канал
	go func() {
		for i := 1; ; i++ {
			// Отправляем значение в канал
			dataChannel <- i
			// Ждем 1 секунду перед отправкой следующего значения
			time.Sleep(time.Second)
		}
	}()

	// Создаем горутину для чтения данных из канала
	go func() {
		for {
			// Читаем данные из канала
			data := <-dataChannel
			// Выводим полученные данные в stdout
			fmt.Println("Received:", data)
		}
	}()

	// Ждем N секунд перед завершением программы
	time.Sleep(time.Duration(N) * time.Second)
	fmt.Println("Program finished.")
}
