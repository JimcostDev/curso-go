package functions

// Las funciones sirven para encapsular código y reutilizarlo

// Algunos puntos importantes adicionales:
// Todas las funciones en Go deben definirse con la palabra clave func
// El nombre de la función debe comenzar con una letra
// Si una función empieza con mayúscula, será exportada (pública)
// Si empieza con minúscula, será privada al paquete

import "fmt"

// Recibe un parámetro y solo realiza una acción (imprimir)
func saludar(nombre string) {
	fmt.Printf("¡Hola %s!\n", nombre)
}

func Simple() {
	fmt.Println("Función simple:")
	saludar("Ronaldo")
	fmt.Println()
}
