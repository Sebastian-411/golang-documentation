package main

import (
    "fmt"
)

func main() {
    // Arrays: una colección de elementos del mismo tipo con un tamaño fijo.
    var arr [5]int // Declaración de un array de enteros con tamaño 5.
    arr[0] = 10
    arr[1] = 20
    arr[2] = 30
    arr[3] = 40
    arr[4] = 50
    fmt.Println("Array:", arr) // Imprimir el array completo.

    // Obtener la longitud del array usando len().
    fmt.Println("Length of array:", len(arr))

    // Slices: una secuencia flexible de elementos de un tipo específico.
    // Los slices se crean a partir de arrays y su tamaño es dinámico.
    var slice []int = arr[1:4] // Crear un slice desde el segundo hasta el cuarto elemento del array.
    fmt.Println("Slice:", slice)

    // Modificar un slice también modifica el array subyacente.
    slice[0] = 100
    fmt.Println("Modified Slice:", slice)
    fmt.Println("Modified Array:", arr) // El array también se modifica.

    // Obtener la longitud y la capacidad de un slice usando len() y cap().
    fmt.Println("Length of slice:", len(slice))
    fmt.Println("Capacity of slice:", cap(slice))

    // Usando make() para crear un slice con tamaño y capacidad específicos.
    newSlice := make([]int, 5, 10) // Crear un slice con longitud 5 y capacidad 10.
    newSlice[0] = 5
    newSlice[1] = 10
    fmt.Println("New Slice with make():", newSlice)

    // Obtener la longitud y la capacidad del nuevo slice.
    fmt.Println("Length of newSlice:", len(newSlice))
    fmt.Println("Capacity of newSlice:", cap(newSlice))

    // A diferencia de los arrays, los slices pueden redimensionarse.
    newSlice = append(newSlice, 15, 20, 25)
    fmt.Println("Resized Slice with append():", newSlice)

    // Obtener la nueva longitud y capacidad después de redimensionar.
    fmt.Println("New length of newSlice:", len(newSlice))
    fmt.Println("New capacity of newSlice:", cap(newSlice))

    // Iterar sobre un slice usando un bucle for.
    fmt.Println("Iterating over newSlice:")
    for i, v := range newSlice {
        fmt.Printf("Index %d: Value %d\n", i, v)
    }

    // Maps: una colección de pares clave-valor.
    var m map[string]int = map[string]int{
        "Alice": 25,
        "Bob":   30,
    }
    fmt.Println("Map:", m)

    // Añadir elementos a un map.
    m["Charlie"] = 35
    fmt.Println("Updated Map:", m)

    // Acceder a elementos de un map.
    age := m["Alice"]
    fmt.Println("Alice's age:", age)

    // Verificar si una clave existe en el map.
    value, exists := m["Bob"]
    if exists {
        fmt.Println("Bob's age:", value)
    } else {
        fmt.Println("Bob is not in the map")
    }

    // Usando make() para crear un map vacío.
    anotherMap := make(map[string]int) // Crear un map vacío.
    anotherMap["David"] = 40
    fmt.Println("Another Map with make():", anotherMap)

    // Iterar sobre un map usando un bucle for.
    fmt.Println("Iterating over map:")
    for key, value := range m {
        fmt.Printf("Key %s: Value %d\n", key, value)
    }
}