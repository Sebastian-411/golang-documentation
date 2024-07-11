package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
	mutex()
	selectChannels()
	
}

func selectChannels(){
    canal1 := make(chan string)
    canal2 := make(chan string)

    go func() {
        time.Sleep(1 * time.Second)
        canal1 <- "Mensaje desde canal 1"
    }()

    go func() {
        time.Sleep(2 * time.Second)
        canal2 <- "Mensaje desde canal 2"
    }()

    for i := 0; i < 2; i++ {
        select {
        case mensaje1 := <-canal1:
            fmt.Println(mensaje1)
        case mensaje2 := <-canal2:
            fmt.Println(mensaje2)
        }
    }
}


func mutex(){
    var mu sync.Mutex
    var contador int
    var wg sync.WaitGroup

    wg.Add(2)

    // Función que incrementa el contador
    incrementar := func(id int) {
        defer wg.Done()
        for i := 0; i < 5; i++ {
            mu.Lock() // Bloquea el mutex
            contador++
            fmt.Printf("Goroutine %d incrementa el contador a: %d\n", id, contador)
            mu.Unlock() // Desbloquea el mutex
            time.Sleep(200 * time.Millisecond)
        }
    }

    // Iniciar dos goroutines que incrementan el contador
    go incrementar(1)
    go incrementar(2)

    wg.Wait() // Espera a que ambas goroutines terminen
    fmt.Printf("Valor final del contador: %d\n", contador)

}