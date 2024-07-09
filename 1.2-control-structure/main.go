package main

import (
	"fmt"
)

func main() {
	// Estructura de control: if-else
	age := 20
	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a minor.")
	}

	// Estructura de control: if-else if-else
	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else if score >= 60 {
		fmt.Println("Grade: D")
	} else {
		fmt.Println("Grade: F")
	}

	// Estructura de control: switch
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}

	// Bucle: for (condicional)
	fmt.Println("For loop (condicional):")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Bucle: for (bucle while)
	fmt.Println("For loop (while):")
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// Bucle: for (infinito)
	fmt.Println("For loop (infinito):")
	k := 0
	for {
		if k >= 5 {
			break // Salir del bucle
		}
		fmt.Println(k)
		k++
	}

	// Bucle: for con range (sobre un slice)
	fmt.Println("For loop con range (slice):")
	numbers := []int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Bucle: for con range (sobre un mapa)
	fmt.Println("For loop con range (map):")
	person := map[string]string{"name": "John", "age": "30"}
	for key, value := range person {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	// Bucle: for con range (sobre una cadena)
	fmt.Println("For loop con range (cadena):")
	for index, char := range "hello" {
		fmt.Printf("Index: %d, Char: %c\n", index, char)
	}

	// Ejemplo de uso de continue en un bucle for
	fmt.Print("Uso de continue en un bucle for: \t")

	// Imprimir números del 0 al 9, excepto el número 5
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue // Saltar la iteración actual si i es igual a 5
		}
		fmt.Print(i)
	}
	fmt.Println("")

	// Ejemplo de uso de break en un bucle for
	fmt.Print("Uso de continue en un bucle for: \t")

	// Imprimir números del 0 al 9, excepto el número 5
	for i := 0; i < 10; i++ {
		if i == 5 {
			break // Salir de la iteración actual si i es igual a 5
		}
		fmt.Print(i)
	}
	fmt.Println("")

}
