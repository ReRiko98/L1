package main

import (
	"fmt"
	"sync"
	"time"
	"context"
)

func workerWithChannel(stop chan struct{}) {
	for {
		select {
		case <-stop:
			fmt.Println("Worker with channel stopped.")
			return
		default:
			// Работа, выполняемая горутиной
			fmt.Println("Working with channel...")
			time.Sleep(time.Second)
		}
	}
}

func workerWithContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker with context stopped.")
			return
		default:
			// Работа, выполняемая горутиной
			fmt.Println("Working with context...")
			time.Sleep(time.Second)
		}
	}
}

func workerWithDone(done chan bool) {
	for {
		select {
		case <-done:
			fmt.Println("Worker with done stopped.")
			return
		default:
			// Работа, выполняемая горутиной
			fmt.Println("Working with done...")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	// 1. Использование канала для управления выполнением горутины
	stopChannel := make(chan struct{})
	go workerWithChannel(stopChannel)
	time.Sleep(3 * time.Second) // Пусть горутина поработает 3 секунды
	close(stopChannel)          // Остановка горутины с помощью закрытия канала

	// 2. Использование контекста (context) для отмены выполнения горутины
	ctx, cancel := context.WithCancel(context.Background())
	go workerWithContext(ctx)
	time.Sleep(3 * time.Second) // Пусть горутина поработает 3 секунды
	cancel()                     // Отмена выполнения горутины с помощью вызова cancel()

	// 3. Использование сигнала завершения (done) для указания завершения работы
	doneChannel := make(chan bool)
	go workerWithDone(doneChannel)
	time.Sleep(3 * time.Second) // Пусть горутина поработает 3 секунды
	doneChannel <- true         // Остановка горутины с помощью отправки сигнала в done канал

	// Ожидание завершения всех горутин
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
	}()
	wg.Wait()
	fmt.Println("Main goroutine finished.")
}
