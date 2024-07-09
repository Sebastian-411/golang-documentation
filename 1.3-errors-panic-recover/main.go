package main

import (
	"errors"
	"fmt"
)

/**
Error en Go: Un valor que indica una condición inesperada o no deseada durante la ejecución del programa, manejado devolviéndolo explícitamente desde funciones.

Panic en Go: Una situación excepcional que detiene abruptamente la ejecución normal del programa debido a una condición crítica e inesperada.

Recover en Go: Una función utilizada para recuperar el control después de un panic, permitiendo la ejecución controlada del programa y el manejo de la situación excepcional.
**/

// Función divide realiza una división entera y maneja errores.
func divide(a, b int) (int, error) {
	// Manejo de errores comunes en Go:
	// - División por cero: Retorna un error personalizado si b es cero.
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

// Función simulatePanic simula un panic.
func simulatePanic() {
	fmt.Println("Simulating a panic")
	// Un panic es una situación excepcional que detiene la ejecución normal del programa.
	panic("panic occurred")
}

func main() {
	// Manejo de errores
	result, err := divide(10, 0)
	// En Go, los errores se manejan devolviendo un valor de error explícito junto con el resultado.
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Manejo de panic con recover
	// Recover se utiliza para manejar panics y permitir que el programa continúe ejecutándose.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	simulatePanic()

	fmt.Println("End of main function") // No ejecutado
}

/*
Comentarios adicionales:


- Defer:
  - La palabra clave defer se utiliza para programar la ejecución de una función hasta que la función que la contiene haya terminado.
  - En este caso, se utiliza para posponer la llamada a recover hasta que el control haya salido del ámbito de main.

- Casos de uso:
  - Manejo de errores: Utilizado en divide para devolver errores explícitos cuando se produce una situación inesperada.
  - Panic: Utilizado en simulatePanic para simular una situación crítica que detiene la ejecución normal.
  - Recover: Utilizado en main para capturar panics y permitir que el programa continúe ejecutándose sin terminar abruptamente.

- Recomendaciones:
  - Usa defer con recover para manejar panics de manera controlada y evitar que el programa se cierre inesperadamente.
  - Utiliza errores explícitos para comunicar problemas y asegurar que los usuarios y otros componentes del sistema puedan manejar adecuadamente las fallas.
*/
