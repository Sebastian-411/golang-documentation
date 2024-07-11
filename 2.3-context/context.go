package main

import (
    "context"
    "fmt"
    "time"
)

// Claves para el contexto
type key string

const requestIDKey key = "requestID"

// Función que realiza una operación simulada que puede ser cancelada
func realizarOperacion(ctx context.Context, nombre string, duracion time.Duration) {
    select {
    case <-time.After(duracion):
        fmt.Printf("Operación %s completada\n", nombre)
    case <-ctx.Done():
        fmt.Printf("Operación %s cancelada: %v\n", nombre, ctx.Err())
    }
}

func main() {
    // Contexto base
    baseCtx := context.Background()

    // Contexto con cancelación
    ctxCancel, cancel := context.WithCancel(baseCtx)
    defer cancel() // Cancelar el contexto al salir de main

    // Contexto con tiempo de espera de 3 segundos
    ctxTimeout, cancelTimeout := context.WithTimeout(baseCtx, 3*time.Second)
    defer cancelTimeout()

    // Contexto con valor
    ctxValue := context.WithValue(baseCtx, requestIDKey, "12345")

    // Iniciar operaciones con diferentes contextos
    go realizarOperacion(ctxCancel, "A", 5*time.Second)
    go realizarOperacion(ctxTimeout, "B", 5*time.Second)
    go realizarOperacion(ctxValue, "C", 5*time.Second)

    // Simular una cancelación después de 2 segundos
    time.Sleep(2 * time.Second)
    cancel() // Cancelar la operación A

    // Esperar un momento para que las operaciones finalicen
    time.Sleep(6 * time.Second)

    // Obtener el valor del contexto
    if requestID, ok := ctxValue.Value(requestIDKey).(string); ok {
        fmt.Printf("Request ID: %s\n", requestID)
    }
}
