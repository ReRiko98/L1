package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	// Создаем канал для передачи данных
	dataChannel := make(chan int)

	// Количество воркеров, которые будут читать данные из канала
	workerCount := 5 // Установите желаемое количество воркеров

	// Создаем WaitGroup для ожидания завершения всех воркеров
	var wg sync.WaitGroup

	// Канал для получения сигнала завершения
	done := make(chan struct{})

	// Создаем горутину для обработки сигнала завершения (Ctrl+C)
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
		<-sig
		close(done)
	}()

	// Запускаем воркеров
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for {
				select {
				case <-done:
					return // Выход из горутины при получении сигнала завершения
				case data := <-dataChannel:
					// Читаем данные из канала и выводим их в stdout
					fmt.Printf("Воркер %d получил данные: %d\n", workerID, data)
				}
			}
		}(i + 1) // Используем i+1 для нумерации воркеров с 1
	}

	// Постоянно записываем данные в канал в главном потоке
	for i := 1; ; i++ {
		select {
		case <-done:
			// Если получен сигнал завершения, закрываем канал данных и ожидаем завершения всех воркеров
			close(dataChannel)
			wg.Wait()
			return
		default:
			// В противном случае, продолжаем записывать данные в канал
			dataChannel <- i
		}
	}
}
