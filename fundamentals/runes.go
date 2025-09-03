package fundamentals

import (
	"fmt"
)

func Runes() {
	// Definimos una cadena de texto que incluye caracteres especiales y un emoji
	palabra := "Go está chévere👍"

	// --- Recorrido usando for y range (¡La forma correcta!) ---
	fmt.Println("Recorrido usando 'range' (¡con runes!):")
	for i, r := range palabra {
		fmt.Printf("Índice: %d, Rune: %c, Valor Unicode: %U\n", i, r, r)
	}

	fmt.Println("\n--- Recorrido usando 'for' simple (¡con bytes!) ---")
	fmt.Println("Esto puede generar caracteres ilegibles para Unicode:")
	for i := 0; i < len(palabra); i++ {
		fmt.Printf("Índice: %d, Byte: %c, Valor byte: %d\n", i, palabra[i], palabra[i])
	}
}
