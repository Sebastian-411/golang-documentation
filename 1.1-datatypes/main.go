package main

import (
	"fmt"
)

func main() {
	// Boolean: representa un valor verdadero (true) o falso (false)
	var boolVar bool = true
	fmt.Println("Boolean:", boolVar)

	// Integer types: representan números enteros de diferentes tamaños
	var intVar int = 42                      // Int: tamaño depende de la arquitectura del sistema (32 o 64 bits)
	var int8Var int8 = 127                   // Int8: entero de 8 bits con signo (-128 a 127)
	var int16Var int16 = 32767               // Int16: entero de 16 bits con signo (-32768 a 32767)
	var int32Var int32 = 2147483647          // Int32: entero de 32 bits con signo (-2147483648 a 2147483647)
	var int64Var int64 = 9223372036854775807 // Int64: entero de 64 bits con signo (-9223372036854775808 a 9223372036854775807)
	fmt.Println("Int:", intVar)
	fmt.Println("Int8:", int8Var)
	fmt.Println("Int16:", int16Var)
	fmt.Println("Int32:", int32Var)
	fmt.Println("Int64:", int64Var)

	// Byte: alias para uint8, representa un byte (0 a 255)
	var byteVar byte = 255
	fmt.Println("Byte:", byteVar)

	// Unsigned integer types: representan números enteros sin signo de diferentes tamaños
	var uintVar uint = 42                       // Uint: tamaño depende de la arquitectura del sistema (32 o 64 bits)
	var uint8Var uint8 = 255                    // Uint8: entero de 8 bits sin signo (0 a 255)
	var uint16Var uint16 = 65535                // Uint16: entero de 16 bits sin signo (0 a 65535)
	var uint32Var uint32 = 4294967295           // Uint32: entero de 32 bits sin signo (0 a 4294967295)
	var uint64Var uint64 = 18446744073709551615 // Uint64: entero de 64 bits sin signo (0 a 18446744073709551615)
	fmt.Println("Uint:", uintVar)
	fmt.Println("Uint8:", uint8Var)
	fmt.Println("Uint16:", uint16Var)
	fmt.Println("Uint32:", uint32Var)
	fmt.Println("Uint64:", uint64Var)

	// Floating point types: representan números con punto decimal de precisión simple y doble
	var float32Var float32 = 3.14              // Float32: precisión simple (aproximadamente 6-7 dígitos decimales)
	var float64Var float64 = 3.141592653589793 // Float64: precisión doble (aproximadamente 15-16 dígitos decimales)
	fmt.Println("Float32:", float32Var)
	fmt.Println("Float64:", float64Var)

	// Rune: alias para int32, representa un código Unicode
	var runeVar rune = 'A'
	fmt.Println("Rune:", runeVar)

	// Complex types: representan números complejos con parte real e imaginaria
	var complex64Var complex64 = complex(1.2, 3.4)   // Complex64: precisión simple para parte real e imaginaria
	var complex128Var complex128 = complex(1.2, 3.4) // Complex128: precisión doble para parte real e imaginaria
	fmt.Println("Complex64:", complex64Var)
	fmt.Println("Complex128:", complex128Var)

	// Uintptr: entero sin signo lo suficientemente grande para contener el valor de un puntero
	var uintptrVar uintptr = 0x12345678
	fmt.Println("Uintptr:", uintptrVar)
}
