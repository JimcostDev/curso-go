package exercises

import (
	"fmt"
)

// print simula la función print de Python, aceptando un número variable de argumentos.
func print(a ...interface{}) {
	fmt.Println(a...)
}

func PrintPy() {
	// Llamadas a la función print con diferentes números y tipos de argumentos.
	print("Hola, mundo!")
	print("El número es:", 42)
	print("El valor de pi es:", 3.14159)
	print("Podemos combinar:", "texto", "y", 123, "con flotantes", 45.67)

	// La función original fmt.Print es en sí misma una función variádica.
	// Nuestra función print simplemente envuelve a fmt.Print.
}
