// Función con parámetros variables (variadic)
package functions

import "fmt"

// Acepta un número variable de argumentos
// Los argumentos se tratan como un slice dentro de la función
// Se indica con ... antes del tipo del parámetro
func sumarTodos(numeros ...int) int {
	total := 0
	for _, num := range numeros {
		total += num
	}
	return total
}

func Variadic() {
	fmt.Println("Función variadic:")
	fmt.Println(sumarTodos(1, 2, 3))       // Suma: 6
	fmt.Println(sumarTodos(5, 10, 15, 20)) // Suma: 50
	fmt.Println()
}
