// Función que retorna un valor
package functions

import "fmt"

// El tipo de retorno se especifica después de los parámetros
func sumar(a, b int) int {
	return a + b
}

func SimpleReturn() {
	fmt.Println("Función de retorno simple:")
	resultado := sumar(5, 3)
	fmt.Printf("La suma es: %d\n", resultado)
	fmt.Println()
}
