package main

import (
    "fmt"
    "sync"
    "time"
)

// Función que cuenta hasta un número dado con un intervalo específico
func contar(wg *sync.WaitGroup, x int, intervalo int) {
    defer wg.Done()
    for i := 0; i <= x; i += intervalo {
        fmt.Printf("Contando %d con intervalo %d\n", i, intervalo)
        time.Sleep(500 * time.Millisecond) // Simular trabajo
    }
}

func comunicar(mensajes chan<- int, wg *sync.WaitGroup, x int) {
    defer wg.Done()
    for i := 0; i <= x; i++ {
        mensajes <- i
    }
    close(mensajes)
}

func main() {
    var wg sync.WaitGroup

    // Número hasta el cual contar
    x := 10

    // Añadir goroutines al WaitGroup
    wg.Add(4)

    // Iniciar goroutines para contar con diferentes intervalos
    go contar(&wg, x, 1)
    go contar(&wg, x, 2)
    go contar(&wg, x, 3)
    go contar(&wg, x, 4)

    // Esperar a que todas las goroutines terminen
    go func() {
        wg.Wait()
        fmt.Println("Ya no cuento más")
    }()

    // Canal de comunicación
    mensajes := make(chan int)

    // Añadir la goroutine de comunicación al WaitGroup
    wg.Add(1)
    go comunicar(mensajes, &wg, x)

    // Leer del canal de comunicación
    go func() {
        for mensaje := range mensajes {
            fmt.Printf("Recibido del canal: %d\n", mensaje)
        }
        fmt.Println("Ya no me comunico más")
    }()

    // Esperar un momento para que todas las goroutines terminen su trabajo
    time.Sleep(10 * time.Second)
}
