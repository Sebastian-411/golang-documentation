# Goroutines y Concurrencia en Go

## Introducción

Las goroutines son la manera en que Go maneja la concurrencia. Son ligeras y se ejecutan en paralelo, permitiendo que múltiples tareas se realicen al mismo tiempo. Las goroutines se utilizan para realizar operaciones que pueden ejecutarse de manera concurrente sin bloquear la ejecución del programa principal.

## Contenidos
- [¿Qué es una Goroutine?](#qué-es-una-goroutine)
- [Cómo usar Goroutines](#cómo-usar-goroutines)
- [Sincronización con WaitGroup](#sincronización-con-waitgroup)
- [Canales para Comunicación](#canales-para-comunicación)
- [Concurrencia en Go](#concurrencia-en-go)
- [Ejemplos Prácticos](#ejemplos-prácticos)

## ¿Qué es una Goroutine?

Una goroutine es una función que se ejecuta concurrentemente con otras funciones. Las goroutines son muy ligeras en comparación con los hilos del sistema operativo, lo que permite que miles de goroutines se ejecuten simultáneamente en un solo programa.

## Cómo usar Goroutines

Para iniciar una goroutine, simplemente utiliza la palabra clave `go` seguida de la llamada a la función que deseas ejecutar de manera concurrente.

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // Iniciar una goroutine
    go decirHola()

    // Mantener el programa principal en ejecución
    time.Sleep(1 * time.Second)
}

func decirHola() {
    fmt.Println("¡Hola desde una goroutine!")
}
```

## Sincronización con WaitGroup

Para esperar a que varias goroutines terminen su ejecución, se puede utilizar `sync.WaitGroup`. Esto es útil cuando necesitas asegurarte de que todas las goroutines hayan terminado antes de que el programa principal continúe.

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup

    // Añadir goroutines al WaitGroup
    wg.Add(2)

    go func() {
        defer wg.Done()
        fmt.Println("Goroutine 1 finalizada")
    }()

    go func() {
        defer wg.Done()
        fmt.Println("Goroutine 2 finalizada")
    }()

    // Esperar a que todas las goroutines terminen
    wg.Wait()
    fmt.Println("Todas las goroutines han finalizado")
}
```

## Canales para Comunicación

Los canales (`chan`) son utilizados para la comunicación segura entre goroutines. Un canal permite enviar y recibir valores entre goroutines, lo que facilita la sincronización y el intercambio de datos.

```go
package main

import "fmt"

func main() {
    mensajes := make(chan string)

    go func() {
        mensajes <- "Mensaje desde la goroutine"
    }()

    mensaje := <-mensajes
    fmt.Println(mensaje)
}
```

En este ejemplo:
- Se crea un canal `mensajes` para transmitir datos de tipo `string`.
- Una goroutine envía un mensaje al canal.
- El programa principal recibe el mensaje del canal y lo imprime.

## Concurrencia en Go

La concurrencia es la capacidad de un programa para gestionar múltiples tareas al mismo tiempo. Go está diseñado con la concurrencia en mente, y las goroutines junto con los canales hacen que la concurrencia sea fácil y eficiente de manejar.

### Ventajas de la Concurrencia en Go

- **Ligereza**: Las goroutines son mucho más ligeras que los hilos del sistema operativo, permitiendo que miles de goroutines se ejecuten en paralelo con un uso mínimo de recursos.
- **Simplicidad**: Las goroutines y los canales proporcionan un modelo simple y claro para la programación concurrente.
- **Escalabilidad**: La concurrencia en Go permite que las aplicaciones escalen fácilmente para manejar cargas de trabajo más grandes y complejas.

### Ejemplo de Concurrencia

```go
package main

import (
    "fmt"
    "time"
)

func contar(id int) {
    for i := 0; i < 5; i++ {
        fmt.Printf("Goroutine %d: %d\n", id, i)
        time.Sleep(500 * time.Millisecond)
    }
}

func main() {
    for i := 1; i <= 3; i++ {
        go contar(i)
    }

    // Mantener el programa principal en ejecución
    time.Sleep(3 * time.Second)
    fmt.Println("Fin del programa")
}
```

## Ejemplos Prácticos

### Ejemplo 1: Contar Números en Paralelo

```go
package main

import (
    "fmt"
    "time"
)

func contar(id int) {
    for i := 0; i < 5; i++ {
        fmt.Printf("Goroutine %d: %d\n", id, i)
        time.Sleep(500 * time.Millisecond)
    }
}

func main() {
    for i := 1; i <= 3; i++ {
        go contar(i)
    }

    // Mantener el programa principal en ejecución
    time.Sleep(3 * time.Second)
    fmt.Println("Fin del programa")
}
```

### Ejemplo 2: Uso de WaitGroup y Canales

```go
package main

import (
    "fmt"
    "sync"
)

func enviarMensaje(mensajes chan<- string, wg *sync.WaitGroup) {
    defer wg.Done()
    mensajes <- "Hola desde la goroutine"
}

func main() {
    var wg sync.WaitGroup
    mensajes := make(chan string, 1)

    wg.Add(1)
    go enviarMensaje(mensajes, &wg)

    wg.Wait()
    close(mensajes)

    for mensaje := range mensajes {
        fmt.Println(mensaje)
    }
}
```

## Conclusión

Las goroutines y los canales son herramientas poderosas para manejar la concurrencia en Go. Con `sync.WaitGroup` y los canales, puedes sincronizar y comunicarte entre goroutines de manera efectiva. Practicar con estos ejemplos te ayudará a entender mejor cómo funcionan las goroutines y cómo puedes utilizarlas en tus propios proyectos.

---

Espero que este README te sea útil para entender y trabajar con goroutines y la concurrencia en Go. Si tienes alguna pregunta o necesitas más ejemplos, ¡no dudes en pedirlo!
```