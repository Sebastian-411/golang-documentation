package main

import "fmt"

func main() {
	var i interface{} = "Hola"

	// Afirmación de tipo para convertir a string
	s, ok := i.(string)
	if ok {
		fmt.Println("Valor de tipo string:", s)
	} else {
		fmt.Println("No se pudo convertir a string")
	}

	// Intentar convertir a int (fallará)
	_, ok = i.(int)
	if !ok {
		fmt.Println("No se pudo convertir a int")
	}

	// Ejemplo de switch con tipos
	determinarTipo(21)
	determinarTipo("Hola")

}

func determinarTipo(i interface{}) {
	switch valor := i.(type) {
	case int:
		fmt.Printf("%d es de tipo entero\n", valor)
	case string:
		fmt.Printf("\"%s\" es de tipo cadena\n", valor)
	default:
		fmt.Printf("Tipo desconocido para valor: %v\n", valor)
	}
}
