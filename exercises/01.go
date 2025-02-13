package exercises

/*
📜 Ejercicio: Registro de Usuario en Go
	Crear un programa en Go que solicite al usuario su nombre, edad,
	altura y una breve descripción sobre sí mismo. Luego, el
	programa mostrará los datos formateados de manera organizada.

📌 Requisitos del ejercicio:
	1. Usar fmt.Scan para capturar el nombre y la edad.
	2. Usar bufio.NewReader para capturar una descripción completa del usuario.
	3. Mostrar los datos usando fmt.Printf con formato.
*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Exercise01() {
	var nombre string
	var edad int
	var altura float64

	// Capturar nombre y edad con fmt.Scan
	fmt.Print("Ingresa tu nombre: ")
	fmt.Scan(&nombre)

	fmt.Print("Ingresa tu edad: ")
	fmt.Scan(&edad)

	fmt.Print("Ingresa tu altura en metros (ej. 1.75): ")
	fmt.Scan(&altura)

	// Limpiar buffer antes de capturar la descripción
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	// Capturar una línea completa con bufio para la descripción
	fmt.Print("Escribe una breve descripción sobre ti: ")
	descripcion, _ := reader.ReadString('\n') // _ es para ignorar el error, ya que no hemos visto como manejar errores
	descripcion = strings.TrimSpace(descripcion)

	// Mostrar los datos usando fmt.Printf con formato
	fmt.Println("\n📌 **Resumen del Usuario:**")
	fmt.Printf("👤 Nombre: %s\n", nombre)
	fmt.Printf("🎂 Edad: %d años\n", edad)
	fmt.Printf("📏 Altura: %.2f metros\n", altura)
	fmt.Printf("📝 Descripción: %s\n", descripcion)
}
