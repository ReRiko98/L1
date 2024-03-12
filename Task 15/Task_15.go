// var justString string
// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100]
// }

// func main() {
//   someFunc()
// }

// В данном фрагменте кода переменной justString присваивается подстрока очень большой строки, возвращаемой функцией createHugeString. 
// Это может привести к выделению дополнительной памяти под всю большую строку в памяти программы, а не только под первые 100 символов, 
// что может быть неэффективным и даже привести к исчерпанию памяти в случае больших значений.

// Чтобы избежать этой проблемы, можно воспользоваться методом копирования строки. Вот исправленный вариант:

package main

//import "fmt"

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func main() {
	someFunc()
}

// Функция для создания большой строки
func createHugeString(size int) string {
	var result string
	for i := 0; i < size; i++ {
		result += "a" // Добавляем символ "a" в строку
	}
	return result
}

// Этот код будет работать правильно и избежит проблемы с выделением памяти под всю большую строку. 
// Он сохраняет только первые 100 символов строки, возвращаемой функцией createHugeString, 
// в переменную justString, используя операцию среза v[:100].