package main

import (
	"fmt"
)

// Human - структура человека
type Human struct {
	Name string
	Age  int
}

// Action - структура действия с человеком, встраивающая Human
type Action struct {
	Human // встраивание структуры Human
}

// Метод структуры Action
func (a *Action) Walk() {
	fmt.Printf("%s is walking.\n", a.Name)
}

// Метод структуры Action
func (a *Action) Talk() {
	fmt.Printf("%s is talking.\n", a.Name)
}

func main() {
	// Создание объекта типа Action
	human := Action{
		Human: Human{Name: "Alice", Age: 30},
	}

	// Использование методов из структуры Action
	human.Walk()
	human.Talk()
}
