package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Примеры переменных различных типов
	var (
		num     = 42
		str     = "Hello, world!"
		boolean = true
		ch      = make(chan int)
	)

	// Преобразуем переменные в интерфейс и помещаем их в срез
	var values = []interface{}{num, str, boolean, ch}

	// Проверяем тип каждой переменной в срезе
	for _, v := range values {
		switch v := v.(type) {
		case int:
			fmt.Printf("%v is of type int\n", v)
		case string:
			fmt.Printf("%v is of type string\n", v)
		case bool:
			fmt.Printf("%v is of type bool\n", v)
		case chan int:
			fmt.Printf("%v is of type channel\n", v)
		default:
			// Если тип переменной не определен, используем reflect.TypeOf для определения
			fmt.Printf("%v is of unknown type: %v\n", v, reflect.TypeOf(v))
		}
	}
}
