# Generics en Go

## Introducción

Los "generics" (o genéricos) en Go permiten escribir funciones, métodos y tipos que pueden manejar datos de diferentes tipos sin especificar explícitamente esos tipos. Esta característica es extremadamente útil para escribir código flexible y reutilizable en Go.

## Contenidos

- [¿Qué son los Generics en Go?](#qué-son-los-generics-en-go)
- [Ejemplos Prácticos](#ejemplos-prácticos)
- [Consideraciones](#consideraciones)
- [Recursos Adicionales](#recursos-adicionales)

## ¿Qué son los Generics en Go?

Los generics en Go permiten escribir funciones y tipos que pueden trabajar con cualquier tipo de datos. Esto se logra mediante la definición de tipos y funciones parametrizados que aceptan tipos de datos como parámetros.

### Ejemplos de Uso

#### Ejemplo de función genérica

```go
package main

import "fmt"

// Función genérica que imprime cualquier tipo de slice
func imprimirSlice[T any](s []T) {
    for _, v := range s {
        fmt.Println(v)
    }
}

func main() {
    // Uso de la función genérica
    ints := []int{1, 2, 3}
    strings := []string{"uno", "dos", "tres"}

    imprimirSlice(ints)
    imprimirSlice(strings)
}
```

En este ejemplo, `imprimirSlice` es una función genérica que puede imprimir cualquier tipo de slice (`[]int`, `[]string`, etc.) sin necesidad de escribir una función por separado para cada tipo.

## Ejemplos Prácticos

Los generics son especialmente útiles para estructuras de datos, algoritmos y funciones que deben ser independientes del tipo de datos que manejan.

## Consideraciones

- Los generics en Go están diseñados para ser simples y eficientes, manteniendo la legibilidad y la seguridad del tipo de datos.
- A partir de Go 1.18, se introdujeron los primeros tipos de datos genéricos, como slices y mapas genéricos.

## Recursos Adicionales

- [Documentación Oficial de Go sobre Generics](https://go.dev/doc/generics)
- [Artículo: Introducing Generics in Go](https://blog.golang.org/generics-next-step)
- [Video: Generics in Go](https://www.youtube.com/watch?v=uzfGZB8eU88)

