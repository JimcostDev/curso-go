package data_structure

import "fmt"

func Slice() {
	// Crear slice de equipos
	equipos := []string{"Barcelona", "Real Madrid", "Juventus", "PSG"}

	// Agregar un elemento al final
	equipos = append(equipos, "Manchester City")
	equipos = append(equipos, "Bayern Munich")

	// Crear un slice a partir de otro
	favoritos := equipos[0:2]
	equipos[0] = "Cali"

	// Imprimir los equipos
	for idx, equipo := range equipos {
		fmt.Println(idx, equipo)
	}

	fmt.Println("Favoritos: ")
	for _, favorito := range favoritos {
		fmt.Println(favorito)
	}

	// longitud y capacidad
	fmt.Printf("Longitud: %d, Capacidad: %d\n", len(equipos), cap(equipos))
	// capacidad es el tamaño del array subyacente
	// capacidad inicial es 4, al agregar un elemento se duplica a 8

	// tipo
	fmt.Printf("El tipo de dato es: %T\n", equipos)

	// Uso de ... para desempaquetar
	// Definimos dos slices
	numeros1 := []int{1, 2, 3}
	numeros2 := []int{4, 5, 6}

	// Usamos `append` con `...` para combinar ambos slices
	resultado := append(numeros1, numeros2...)
	// ... expande los elementos del slice resultante como
	// argumentos individuales para la función append, Esto es necesario porque append
	// espera una lista de elementos, no otro slice.

	// Mostramos el resultado
	fmt.Println(resultado) // Salida: [1 2 3 4 5 6]
}
