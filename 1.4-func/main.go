package main

import (
    "fmt"
)

// Función normal: una función básica que toma dos enteros y retorna su suma.
func add(a, b int) int {
    return a + b
}

// Función con múltiples retornos: una función que retorna dos valores.
func divideAndRemainder(a, b int) (int, int) {
    return a / b, a % b
}

// Función con punteros: una función que toma punteros como argumentos.
func increment(a *int) {
    *a++
}

// Función lambda (anónima): una función sin nombre definida en línea.
var multiply = func(a, b int) int {
    return a * b
}

// Callback: una función que toma otra función como argumento.
func executeCallback(a, b int, callback func(int, int) int) int {
    return callback(a, b)
}

func main() {
    // Llamando a una función normal.
    // Función normal: Una función que toma argumentos y retorna un resultado. Utilizada para operaciones básicas.
    sum := add(3, 4)
    fmt.Println("Sum:", sum)

    // Llamando a una función con múltiples retornos.
    // Función con múltiples retornos: Permite devolver más de un valor, útil para funciones que necesitan retornar un resultado y un estado.
    quotient, remainder := divideAndRemainder(10, 3)
    fmt.Println("Quotient:", quotient, "Remainder:", remainder)

    // Llamando a una función con punteros.
    // Función con punteros: Manipula directamente los valores de las variables a través de sus direcciones en memoria.
    value := 5
    increment(&value)
    fmt.Println("Incremented Value:", value)

    // Llamando a una función lambda (anónima).
    // Función lambda: Función sin nombre que puede ser asignada a una variable o pasada como argumento.
    product := multiply(3, 4)
    fmt.Println("Product:", product)

    // Llamando a una función callback.
    // Callback: Función que toma otra función como argumento, permitiendo la ejecución dinámica de la función pasada.
    result := executeCallback(5, 6, add)
    fmt.Println("Callback Result:", result)
}
