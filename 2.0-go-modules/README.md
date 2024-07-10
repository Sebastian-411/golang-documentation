#### Go Modules

Los Go modules son una característica introducida en Go 1.11 para gestionar dependencias y versiones de paquetes de manera más estructurada y reproducible.

### ¿Qué son los Go modules?

Los Go modules permiten a los desarrolladores especificar y gestionar las dependencias de sus proyectos Go de manera explícita. Esto facilita la creación de entornos de desarrollo consistentes y reproducibles, independientemente de la ubicación del código fuente.

### Uso básico de Go modules

Para utilizar Go modules en un proyecto, sigue estos pasos:

1. **Inicialización del módulo:**

   Para inicializar un módulo en tu proyecto, ejecuta el siguiente comando en el directorio raíz del proyecto:

   ```bash
   go mod init nombre-del-modulo
   ```

   Esto creará un archivo `go.mod` en el directorio actual que especifica el nombre del módulo y las dependencias directas del proyecto.

2. **Agregar dependencias:**

   Puedes agregar nuevas dependencias a tu proyecto utilizando `go get`. Por ejemplo, para agregar la dependencia del paquete `github.com/ejemplo/paquete`, ejecuta:

   ```bash
   go get github.com/ejemplo/paquete
   ```

   Esto descargará el paquete y actualizará automáticamente el archivo `go.mod` con la información de la nueva dependencia.

3. **Versiones y actualizaciones:**

   Go modules gestionará automáticamente las versiones de las dependencias según las reglas de versionado semántico (semver). Puedes actualizar las dependencias a las versiones más recientes utilizando `go get -u`.

### Ejemplo de código

A continuación, se muestra un ejemplo de cómo se vería el archivo `main.go` utilizando un módulo y una dependencia:

```go
package main

import (
    "fmt"
    "github.com/ejemplo/paquete"
)

func main() {
    fmt.Println("Ejemplo de uso de un paquete con Go modules")

    // Utilizando una función del paquete importado
    paquete.MiFuncion()
}
```

### Recursos adicionales

- [Documentación oficial de Go modules](https://golang.org/doc/go1.11#modules)
- [Artículo sobre Go modules en la Wiki de Go](https://github.com/golang/go/wiki/Modules)