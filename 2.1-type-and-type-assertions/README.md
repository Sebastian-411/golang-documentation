### Tipos (`types`) en Go

En Go, los tipos definen las propiedades y el comportamiento de las variables y valores en tu programa. Aquí tienes algunos tipos básicos en Go:

```go
package main

import "fmt"

func main() {
    // Tipos básicos
    var entero int
    var flotante float64
    var cadena string
    var booleano bool

    entero = 42
    flotante = 3.14
    cadena = "Hola, mundo!"
    booleano = true

    fmt.Println(entero, flotante, cadena, booleano)
}
```

### Type Assertions (`type assertions`) en Go

Las aserciones de tipo en Go te permiten comprobar y convertir interfaces a tipos subyacentes. Aquí tienes un ejemplo de cómo usarlas:

```go
package main

import (
    "fmt"
)

func main() {
    var i interface{} = "Hola"

    // Aserción de tipo
    cadena, ok := i.(string)
    if ok {
        fmt.Println("La variable i es una cadena:", cadena)
    } else {
        fmt.Println("La variable i no es una cadena")
    }
}
```

### Switches en Go

Los switches en Go permiten evaluar múltiples condiciones de manera concisa. Puedes usarlos con tipos, valores, y expresiones. Aquí tienes un ejemplo:

```go
package main

import "fmt"

func main() {
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
```

### Conclusión

Estos ejemplos te introducen a los tipos en Go, cómo realizar aserciones de tipo y cómo usar switches para manejar diferentes casos. Practicar estos conceptos te ayudará a familiarizarte con la sintaxis y el comportamiento de Go en la manipulación de tipos y estructuras de control.
