# Go Scheduler

## Introducción

El Go Scheduler es una parte fundamental del tiempo de ejecución de Go (Go runtime). Se encarga de gestionar la ejecución concurrente de goroutines. Entender cómo funciona el Go Scheduler te ayudará a escribir programas concurrentes más eficientes y a diagnosticar problemas de rendimiento.

## Contenidos

- [¿Qué es el Go Scheduler?](#qué-es-el-go-scheduler)
- [Cómo funciona el Go Scheduler](#cómo-funciona-el-go-scheduler)
- [Configuración del Scheduler](#configuración-del-scheduler)
- [Ejemplo Práctico](#ejemplo-práctico)
- [Recursos Adicionales](#recursos-adicionales)

## ¿Qué es el Go Scheduler?

El Go Scheduler es un componente del tiempo de ejecución de Go que gestiona la ejecución de goroutines, que son unidades ligeras de ejecución. A diferencia de los hilos del sistema operativo, las goroutines son más ligeras y permiten la creación de miles de ellas con un costo de memoria y procesamiento mínimo.

## Cómo funciona el Go Scheduler

El Go Scheduler utiliza un modelo de concurrencia basado en **M:N**:
- **M**: Número de hilos del sistema operativo (OS threads).
- **N**: Número de goroutines.

### Componentes Clave

- **Goroutines (`G`)**: Unidades ligeras de ejecución que representan la lógica de un programa concurrente.
- **Procesos (`P`)**: Estructuras que contienen las colas de goroutines listas para ejecutar y están asociados a hilos del sistema operativo (`M`).
- **Hilos del Sistema Operativo (`M`)**: Hilos del sistema operativo donde se ejecutan las goroutines.

El Scheduler se asegura de que las goroutines se distribuyan eficientemente entre los hilos del sistema operativo, maximizando el uso de recursos y minimizando el tiempo de inactividad.

### Modelo de Ejecución

1. **Creación de Goroutines**: Cuando se crea una nueva goroutine, se asigna a un `P`.
2. **Ejecutar Goroutines**: El `P` selecciona una goroutine de su cola de goroutines listas y la ejecuta en un `M`.
3. **Bloqueo y Desbloqueo**: Si una goroutine se bloquea (por ejemplo, esperando I/O), el `M` puede liberar el `P` y buscar otra goroutine para ejecutar.
4. **Stealing**: Si un `P` se queda sin goroutines para ejecutar, puede "robar" goroutines de otro `P` para equilibrar la carga.

## Configuración del Scheduler

El comportamiento del Scheduler puede configurarse usando la variable de entorno `GOMAXPROCS`, que controla el número máximo de hilos del sistema operativo que puede utilizar el runtime de Go simultáneamente.

### Uso de `GOMAXPROCS`

```go
package main

import (
    "fmt"
    "runtime"
)

func main() {
    // Establecer GOMAXPROCS a 2
    runtime.GOMAXPROCS(2)
    fmt.Println("Número de CPUs:", runtime.NumCPU())
    fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))
}
```

## Ejemplo Práctico

Aquí hay un ejemplo práctico que demuestra cómo el Scheduler maneja múltiples goroutines.

```go
package main

import (
    "fmt"
    "runtime"
    "sync"
)

func main() {
    runtime.GOMAXPROCS(2) // Configura GOMAXPROCS para utilizar 2 hilos del sistema operativo

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
```

### Explicación del Código

1. **Configuración de `GOMAXPROCS`**: Configura el número de hilos del sistema operativo a 2.
2. **Goroutines**: Crea tres goroutines, cada una de las cuales ejecuta un bucle de 5 iteraciones.
3. **WaitGroup**: Utiliza un `sync.WaitGroup` para esperar a que todas las goroutines finalicen antes de imprimir "Finalizado".

## Recursos Adicionales

- [Documentación Oficial de Go](https://golang.org/doc/)
- [The Go Blog - Concurrency](https://blog.golang.org/concurrency-is-not-parallelism)
- [Video: Inside the Go Scheduler](https://www.youtube.com/watch?v=5zXAHh5tJqQ)

