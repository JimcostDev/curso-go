package exercises

import (
	"fmt"
)

/*
   ---- CONJETURA DE GOLDBACH ----
   Todo número par mayor que 2 puede expresarse como suma de dos números primos.

   Ejemplo para 14: [3 + 11 = 14], [7 + 7 = 14]

   El programa muestra todas las combinaciones únicas de primos que suman el número dado.
*/

func Goldbach() {
	var numeroPar int
	fmt.Println("Ingrese un número PAR mayor que 2 (ej: 4, 14, 28, etc):")
	fmt.Scanln(&numeroPar)

	// Validación básica del input
	if numeroPar <= 2 || numeroPar%2 != 0 {
		fmt.Printf("\n❌ %d no es un número par válido. Debe ser par y mayor que 2\n", numeroPar)
		return
	}

	mitad := numeroPar / 2
	var parejas [][2]int // Slice para almacenar las parejas de primos

	// Buscar parejas solo hasta la mitad del número para evitar duplicados
	for a := 2; a <= mitad; a++ {
		if EsPrimo(a) {
			b := numeroPar - a // 4 - 7
			if EsPrimo(b) {
				parejas = append(parejas, [2]int{a, b})
			}
		}
	}

	// Mostrar resultados
	fmt.Printf("\n🔍 Resultados para %d:\n", numeroPar)
	if len(parejas) > 0 {
		fmt.Printf("Se encontraron %d combinación(es):\n", len(parejas))
		for _, pareja := range parejas {
			fmt.Printf("- %d + %d = %d\n", pareja[0], pareja[1], numeroPar)
		}
	} else {
		fmt.Printf("⚠️  No se encontraron parejas de primos (¡esto contradice la conjetura!)\n")
	}
}
