package main

import (
    "fmt"
)

// Definición de un struct llamado Person
type Person struct {
    Name    string
    Age     int
    Email   string // Campo opcional
    Address Address // Embedding de otro struct
}

// Definición de otro struct para la dirección
type Address struct {
    City    string
    Country string
}

// Método para el struct Person
func (p Person) Greet() {
    fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

// Interface para objetos que pueden describirse
type Describable interface {
    Describe() string
}

// Implementación del método Describe para Person
func (p Person) Describe() string {
    if p.Email != "" {
        return fmt.Sprintf("%s is %d years old, living in %s, %s. Contact: %s", p.Name, p.Age, p.Address.City, p.Address.Country, p.Email)
    }
    return fmt.Sprintf("%s is %d years old, living in %s, %s.", p.Name, p.Age, p.Address.City, p.Address.Country)
}

func main() {
    // Crear una instancia de Person sin inicializar Email
    person := Person{
        Name: "Alice",
        Age:  30,
        Address: Address{
            City:    "New York",
            Country: "USA",
        },
    }

    // Llamar al método Greet
    person.Greet()

    // Llamar al método Describe a través de la interfaz
    var d Describable
    d = person
    fmt.Println(d.Describe())

    // Crear otra instancia de Person con Email
    personWithEmail := Person{
        Name:  "Bob",
        Age:   25,
        Email: "bob@example.com",
        Address: Address{
            City:    "San Francisco",
            Country: "USA",
        },
    }

    // Llamar al método Describe para personWithEmail
    fmt.Println(personWithEmail.Describe())
}
