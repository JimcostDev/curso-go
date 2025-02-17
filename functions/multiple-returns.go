// Función que retorna múltiples valores
package functions

import "fmt"

// Go permite retornar múltiples valores,   estos se especifican entre paréntesis
// Los valores se pueden asignar a múltiples variables
func operacionesMatematicas(a, b int) (int, int, int) {
	suma := a + b
	resta := a - b
	multiplicacion := a * b
	return suma, resta, multiplicacion
}

func MultipleReturns() {
	fmt.Println("Función de múltiples retornos:")
	sum, rest, mult := operacionesMatematicas(10, 5)
	fmt.Printf("Suma: %d, Resta: %d, Multiplicación: %d\n", sum, rest, mult)
	fmt.Println()
}
