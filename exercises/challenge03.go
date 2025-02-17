package exercises

// El objetivo de este ejercicio es practicar el uso de bucles (for),
// estructuras condicionales (if-else) y
// la generación de números aleatorios con rand. Debes
// implementar un juego donde el usuario intente adivinar
// un número entre 1 y 100 con un número limitado de intentos.
// ¡Inténtalo antes de ver la solución!

import (
	"fmt"
	"math/rand"
)

func Challenge03() {
	// Genera número aleatorio
	numeroSecreto := rand.Intn(100) + 1

	// Configuración del juego
	const intentosMaximos = 5
	fmt.Println("¡Bienvenido al juego de adivinanzas!")
	fmt.Println("Adivina un número entre 1 y 100")

	// Bucle principal del juego
	for intento := 1; intento <= intentosMaximos; intento++ {
		fmt.Printf("\nIntento %d/%d: ", intento, intentosMaximos)

		var numeroUsuario int
		fmt.Scan(&numeroUsuario)

		// Comprueba la respuesta
		if numeroUsuario == numeroSecreto {
			fmt.Println("¡Felicidades! ¡Has adivinado el número!")
			return
		}

		// Da pistas usando if-else
		if numeroUsuario < numeroSecreto {
			fmt.Println("Pista: El número es mayor")
		} else {
			fmt.Println("Pista: El número es menor")
		}

		// Mensaje de último intento
		if intento == intentosMaximos {
			fmt.Printf("\nGame Over. El número era: %d\n", numeroSecreto)
		}
	}
}
