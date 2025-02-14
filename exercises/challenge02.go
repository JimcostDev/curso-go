package exercises

// El objetivo de este ejercicio es practicar tipos personalizados (type) y genéricos (generics).
// Debes crear una función genérica que calcule el promedio de dos valores y
// usar un tipo personalizado para representar calificaciones.
// ¡Inténtalo antes de ver la solución!

import (
	"fmt"
)

// Calificacion es un tipo personalizado para representar calificaciones
type Calificacion float64

// Promedio es una función genérica para calcular el promedio de dos valores
func Promedio[T float64 | Calificacion](valor1, valor2 T) T {
	return (valor1 + valor2) / 2
}

func Challenge02() {
	// Usar el tipo personalizado Calificacion
	var math Calificacion = 9.5
	var ingles Calificacion = 7.8

	fmt.Printf("La calificación de matemáticas es: %v\n", math)
	fmt.Printf("La calificación de inglés es: %v\n", ingles)

	// Promedio de calificaciones
	promedio := Promedio(math, ingles)
	fmt.Printf("El promedio de las calificaciones es: %v\n", promedio)

	// Promedio de exámenes
	examen1 := 7.5
	examen2 := 8.3
	promedioExamenes := Promedio(examen1, examen2)
	fmt.Printf("El promedio de los exámenes es: %v\n", promedioExamenes)
}
