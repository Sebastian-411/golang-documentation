package main

import (
    "fmt"
    "runtime"
    "sync"
)

func main() {
    runtime.GOMAXPROCS(5) 

    var wg sync.WaitGroup
    wg.Add(3)

    go func() {
        defer wg.Done()
        for i := 0; i < 5; i++ {
            fmt.Println("Goroutine 1 - Iteración", i)
        }
    }()

    go func() {
        defer wg.Done()
        for i := 0; i < 5; i++ {
            fmt.Println("Goroutine 2 - Iteración", i)
        }
    }()

    go func() {
        defer wg.Done()
        for i := 0; i < 5; i++ {
            fmt.Println("Goroutine 3 - Iteración", i)
        }
    }()

    wg.Wait()
    fmt.Println("Finalizado")
}
