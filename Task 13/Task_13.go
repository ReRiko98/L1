package main

import "fmt"

func main() {
    // Исходные числа
    a := 5
    b := 10
    
    fmt.Println("Before swapping:")
    fmt.Println("a =", a)
    fmt.Println("b =", b)
    
    // Меняем местами числа без использования временной переменной
    a = a + b
    b = a - b
    a = a - b
    
    fmt.Println("After swapping:")
    fmt.Println("a =", a)
    fmt.Println("b =", b)
}
