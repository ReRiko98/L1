package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создаем мьютекс для синхронизации доступа к map
	var mu sync.Mutex

	// Создаем map для хранения данных
	data := make(map[string]int)

	// Создаем WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Функция для конкурентной записи данных в map
	writeToMap := func(key string, value int) {
		// Блокируем мьютекс перед записью данных в map
		mu.Lock()
		// Задержка для имитации работы
		defer mu.Unlock()
		defer wg.Done()
		data[key] = value
		fmt.Printf("Data with key '%s' and value %d written to map\n", key, value)
	}

	// Запускаем несколько горутин для конкурентной записи данных в map
	wg.Add(3)
	go writeToMap("key1", 10)
	go writeToMap("key2", 20)
	go writeToMap("key3", 30)

	// Ожидаем завершения всех горутин
	wg.Wait()

	// Выводим содержимое map после всех записей
	fmt.Println("Map contents after writes:", data)
}
