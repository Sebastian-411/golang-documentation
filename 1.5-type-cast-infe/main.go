package main

import (
	"fmt"
)

func main() {
	/**
	  - Type Casting en Go se utiliza para convertir explícitamente un valor de un tipo a otro.
	**/
	var i int = 42
	// Convertimos el entero i a un flotante.
	var f float64 = float64(i)
	fmt.Println("Integer:", i, "Float:", f)

	// Type Inference: Inferencia automática del tipo de dato basado en el valor asignado.
	// Go infiere que j es un entero basado en el valor asignado.
	j := 43
	fmt.Printf("Type of j: %T, Value of j: %d\n", j, j)

	// Type Casting entre tipos numéricos.
	var x int64 = 10
	var y int32 = int32(x) // Convertimos int64 a int32.
	fmt.Printf("Original int64: %d, Converted int32: %d\n", x, y)

	// Type Casting entre tipos compatibles.
	var s string = "100"
	var n int
	// Convertimos la cadena s a un entero.
	fmt.Sscanf(s, "%d", &n)
	fmt.Printf("String: %s, Converted int: %d\n", s, n)

	// Ejemplo adicional de Type Inference.
	// Go infiere que k es un flotante basado en el valor asignado.
	k := 42.5
	fmt.Printf("Type of k: %T, Value of k: %f\n", k, k)
}

