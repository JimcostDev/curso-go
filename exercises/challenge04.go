package exercises

// Un número primo es un número natural mayor que 1 que tiene exactamente dos divisores positivos distintos: él mismo y el 1.
// Por ejemplo: 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, etc.

// Para determinar si un número es primo, puedes seguir estos pasos:
// 1. Verifica si el número es menor o igual a 1: Si es así, no es primo.
// 2. El número 2 es el único número par primo.
// 3. Descarta los números pares mayores que 2: Cualquier número par mayor que 2 no es primo.
// 4. Aplica la prueba de divisibilidad hasta la raíz cuadrada del número:
//   - Calcula la raíz cuadrada del número.
//   - Verifica si el número es divisible por algún número primo menor o igual a esa raíz cuadrada.
//   - Si es divisible por alguno de ellos, no es primo; de lo contrario, es primo.
// Por ejemplo, para verificar si 29 es primo:
// - La raíz cuadrada de 29 es aproximadamente 5.39.
// - Los números primos menores o iguales a 5 son 2, 3 y 5.
// - 29 no es divisible por 2, 3 ni 5.

// Crear 2 funciones:
// - 1. la función debe recibir un número entero y devolver true si es primo y false si no lo es.
// - 2. la función debe recibir un número entero y imprimir todos los números primos hasta ese número.
// ¡Inténtalo antes de ver la solución!

import (
	"fmt"
	"math"
)

// Verificar si un número es primo.
func EsPrimo(n int) bool {
	// Los números menores o iguales a 1 no son primos
	if n <= 1 {
		return false
	}
	// 2 es el único número par primo
	if n == 2 {
		return true
	}
	// Los números pares mayores que 2 no son primos
	if n%2 == 0 {
		return false
	}
	// Solo se prueban divisores impares desde 3 hasta la raíz cuadrada de n
	limite := int(math.Sqrt(float64(n)))
	for i := 3; i <= limite; i += 2 { // divisores impares, por eso se incrementa en 2
		if n%i == 0 {
			return false
		}
	}
	return true // Si no se encontró ningún divisor, el número es primo
}

// Generar números primos hasta n.
func generarPrimos(n int) []int {
	primos := []int{}
	for i := 2; i <= n; i++ {
		if EsPrimo(i) {
			primos = append(primos, i)
		}
	}
	return primos
}

func Challenge04() {
	hasta := 50
	primos := generarPrimos(hasta)
	fmt.Printf("Números primos hasta %d: %v\n", hasta, primos)
}
