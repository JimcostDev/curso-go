package fundamentals

// Principales especificadores de formato en Go:
// %s: Para cadenas de texto.
// %d: Para enteros decimales.
// %f: Para números de punto flotante.
// %.nf: Para números de punto flotante con n decimales. ejem: %.2f
// %t: Para valores booleanos.
// %c: Para caracteres individuales.
// %b: Para representación binaria de un entero.
// %e: Para notación científica.
// %x: Para representación hexadecimal.
// %T: Para mostrar el tipo de una variable.
// %v: Para imprimir el valor de una variable, independientemente de su tipo.

import "fmt"

// formateo de texto en Go
func Formatted() {
	var nombre string = "Ronaldo"
	var edad int = 25
	var altura float64 = 1.84
	var esEstudiante bool = true
	var caracter rune = 'R'

	fmt.Println("Ejemplo de uso de auxiliares de formato:")
	fmt.Printf("Nombre: %s\n", nombre)
	fmt.Printf("Nombre: %v\n", nombre)
	fmt.Printf("Edad: %d años\n", edad)
	fmt.Printf("Altura: %.2f metros\n", altura)
	fmt.Printf("Es estudiante: %t\n", esEstudiante)
	fmt.Printf("Primer letra del nombre: %c\n", caracter)
	fmt.Printf("Edad en binario: %b\n", edad)
	fmt.Printf("Altura en notación científica: %e\n", altura)
	fmt.Printf("Edad en hexadecimal: %x\n", edad)
	fmt.Printf("Tipo de la variable 'nombre': %T\n", nombre)
	fmt.Println()

	// alinear la salida
	fmt.Printf("%-10s | %10s\n", "Nombre", "Edad")
	fmt.Printf("%-10s | %10d\n", "Carlos", 25)
	fmt.Printf("%-10s | %10d\n", "Sofía", 30)
	fmt.Println()

	// bases numéricas
	num := 42
	fmt.Printf("Decimal: %d\n", num)
	fmt.Printf("Binario: %b\n", num)
	fmt.Printf("Octal: %o\n", num)
	fmt.Printf("Hexadecimal: %x\n", num)

	// guardar la salida en una variable
	s := fmt.Sprintf("Nombre: %s, Edad: %d, Altura: %.2f", nombre, edad, altura)
	fmt.Println(s)
}
