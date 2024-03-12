package main

import (
	"fmt"
	"time"
)

// Функция Sleep приостанавливает выполнение программы на duration миллисекунд
func Sleep(duration time.Duration) {
	time.Sleep(duration * time.Millisecond)
}

func main() {
	fmt.Println("Начало работы программы")

	// Вызываем функцию Sleep на 2 секунды (2000 миллисекунд)
	Sleep(2000)

	fmt.Println("Программа завершена")
}
