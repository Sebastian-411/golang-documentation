package main

import "fmt"

// Función genérica que imprime cualquier tipo de slice
func imprimirSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func main() {
	ints := []int{1, 2, 3}
	strings := []string{"uno", "dos", "tres"}

	imprimirSlice(ints)
	imprimirSlice(strings)
}
