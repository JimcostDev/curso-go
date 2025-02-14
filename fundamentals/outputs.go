package fundamentals

import "fmt"

// Ejemplos con Print, Println y Printf
func prints() {
	//  Imprime los argumentos que se le pasan sin añadir espacios entre ellos ni un salto de línea al final.
	fmt.Print("Hola", " ", "Mundo") // Salida: Hola Mundo

	//Imprime los argumentos que se le pasan, añadiendo un espacio entre ellos y un salto de línea al final.
	fmt.Println("Hola", "Mundo") // Salida: Hola Mundo\n

	// Imprime los argumentos según un formato especificado. Es similar a printf en C. Permite formatear cadenas, números, etc.
	fmt.Printf("Hola %s, tienes %d años\n", "Mundo", 2025) // Salida: Hola Mundo, tienes 2025 años\n
}

// Ejemplos con Sprint, Sprintln y Sprintf
func sprints() {
	// Devuelve una cadena con los argumentos concatenados, sin añadir espacios ni saltos de línea.
	s := fmt.Sprint("Hola", " ", "Mundo") // s = "Hola Mundo"
	fmt.Print(s)

	// Devuelve una cadena con los argumentos concatenados, añadiendo un espacio entre ellos y un salto de línea al final.
	x := fmt.Sprintln("Hola", "Mundo") // x = "Hola Mundo\n"
	fmt.Print(x)

	// Devuelve una cadena formateada según el formato especificado.
	t := fmt.Sprintf("Hola %s, tienes %d años", "Mundo", 2025) // t = "Hola Mundo, tienes 2025 años"
	fmt.Print(t)
}

func Outputs() {
	prints()
	fmt.Println("-------------------")
	sprints()

	// tambien existe Fprint, Fprintln y Fprintf, que permiten escribir en un archivo o en un buffer.
	// Su uso es similar a Print, Println y Printf.
}
