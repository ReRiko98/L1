package main

import (
	"fmt"
	"math"
)

// Структура Point представляет точку в двумерном пространстве
type Point struct {
	x, y float64
}

// Конструктор для создания новой точки
func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

// Метод для вычисления расстояния между текущей точкой и другой точкой
func (p *Point) DistanceTo(other *Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	// Создаем две точки
	point1 := NewPoint(1, 2)
	point2 := NewPoint(4, 6)

	// Находим расстояние между точками
	distance := point1.DistanceTo(point2)

	// Выводим результат
	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
