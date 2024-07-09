package main

import (
    "fmt"
)

func main() {
    // Declaración básica de variables con valor inicial
    var name string = "Alice" 
    fmt.Println("Name:", name)

    // Declaración sin valor inicial (valor cero predeterminado)
    var age int 
    fmt.Println("Age:", age) // age será 0 por defecto

    // Declaración múltiple en una línea
    var length, width int = 5, 10
    fmt.Println("Length:", length, "Width:", width)

    // Uso de la sintaxis corta para declarar e inicializar
    height := 15 
    fmt.Println("Height:", height)

    // Declaración de constantes
    const pi = 3.14
    fmt.Println("Pi:", pi)

    // Variables en bloques
    var (
        city   string = "Paris"
        country string = "France"
    )
    fmt.Println("City:", city, "Country:", country)

    // Declaración y uso de punteros
    var ptr *int
    number := 42
    ptr = &number // ptr apunta a la dirección de memoria de number
    fmt.Println("Pointer value:", *ptr) // Imprime el valor al que apunta ptr (42)

    // Variables sin usar
    // var unusedVar int // Esto causará un error de compilación si no se usa

    // Uso de variables en diferentes contextos
    // Ejemplo: contadores en bucles
    for i := 0; i < 5; i++ {
        fmt.Println("Counter:", i)
    }

    // Ejemplo: capturar entradas del usuario
    var userInput string
    fmt.Print("Enter something: ")
    fmt.Scanln(&userInput)
    fmt.Println("You entered:", userInput)

    // Ejemplo: uso en funciones
    result := add(3, 4)
    fmt.Println("Addition Result:", result)
}

// Función para sumar dos números
func add(a int, b int) int {
    return a + b
}
