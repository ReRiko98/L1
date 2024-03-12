package main

import (
	"fmt"
)

// Square - структура, представляющая квадрат
type Square struct {
	side float64
}

// Area - метод для вычисления площади квадрата
func (s *Square) Area() float64 {
	return s.side * s.side
}

// Shape - интерфейс, определяющий метод GetArea()
type Shape interface {
	GetArea() float64
}

// SquareAdapter - адаптер для квадрата, преобразующий его в объект, удовлетворяющий интерфейсу Shape
type SquareAdapter struct {
	square *Square
}

// GetArea - метод для вычисления площади квадрата через адаптер
func (sa *SquareAdapter) GetArea() float64 {
	return sa.square.Area()
}

func main() {
	// Создаем объект квадрата
	square := &Square{side: 5}

	// Создаем адаптер для квадрата
	adapter := &SquareAdapter{square: square}

	// Используем адаптер для вычисления площади квадрата через интерфейс Shape
	area := adapter.GetArea()

	// Выводим результат
	fmt.Println("Площадь квадрата:", area)
}
