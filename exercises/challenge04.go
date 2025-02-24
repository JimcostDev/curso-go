package exercises

// Un número primo es un número entero mayor que 1 que solo
// es divisible por 1 y por sí mismo. Por ejemplo: 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, etc.
// Crear 2 funciones:
// - 1. la función debe recibir un número entero y devolver true si es primo y false si no lo es.
// - 2. la función debe recibir un número entero y imprimir todos los números primos hasta ese número.
// ¡Inténtalo antes de ver la solución!

import "fmt"

// Verificar si un número es primo de forma sencilla.
func esPrimo(n int) bool {
	if n < 2 {
		return false // los numeros menores que 2 no son primos
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false // si encuentra un divisor, n no es primo
		}
	}
	return true // si no encuentra divisores, n es primo
}

// Generar números primos hasta n.
func generarPrimos(n int) {
	for i := 2; i <= n; i++ {
		if esPrimo(i) {
			fmt.Println(i)
		}
	}
}

func Challenge04() {
	generarPrimos(50)
}
