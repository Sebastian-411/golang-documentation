package main

import "fmt"

// Declaración de una interfaz llamada Describible
type Describible interface {
    Describe() string
}

// Definición de un struct llamado Persona
type Persona struct {
    Nombre string
    Edad   int
}

// Método para implementar la interfaz Describible
func (p Persona) Describe() string {
    return fmt.Sprintf("%s tiene %d años", p.Nombre, p.Edad)
}

// Función que acepta cualquier tipo que implemente Describible
func ImprimirDescripción(d Describible) {
    fmt.Println(d.Describe())
}


func main() {
    // Crear una instancia de Persona
    persona := Persona{
        Nombre: "Juan",
        Edad:   30,
    }

    // Llamar al método Describe a través de la interfaz
    var d Describible
    d = persona
    fmt.Println(d.Describe()) // Salida: Juan tiene 30 años

	persona2 := Persona{
        Nombre: "María",
        Edad:   25,
    }

    ImprimirDescripción(persona2) // Se puede pasar Persona porque implementa Describible
}
