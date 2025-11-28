// Tienes una función que debería sumar los elementos de un slice de números.
package exercises

import "fmt"

// func calcularSuma(numeros []int) int {
// 	var total int
// 	for i := 0; i < len(numeros); i++ {
// 		if numeros[i] > 0 { // El Bug: Solo suma positivos, pero la descripción  dice "sumar los elementos".
// 			total = total + numeros[i]
// 		}
// 	}
// 	return total
// }

//	func Refactoring() {
//		datos := []int{10, 5, -3, 8}
//		resultado := calcularSuma(datos)
//		fmt.Println("La suma total es:", resultado)
//		// Debería ser 20 (10+5-3+8), pero el código devuelve 23. ¡Hay un bug!
//	}

// Refactorización 1: Usar un rango 'range' en lugar de un bucle for tradicional.
// Es más idiomático y seguro en Go.
func calcularSuma(numeros []int) int {
	var total int

	// El 'range' de Go hace el código más limpio. El "_" ignora el índice.
	for _, valor := range numeros {
		total += valor // Refactorización 2: Usar el operador compuesto +=
	}

	return total
}

func Refactoring() {
	datos := []int{10, 5, -3, 8}
	// El resultado ahora será 20. ¡El bug ha sido corregido!
	resultado := calcularSuma(datos)
	fmt.Println("La suma total es:", resultado)
}
