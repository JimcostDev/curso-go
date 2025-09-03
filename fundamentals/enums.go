package fundamentals

import "fmt"

// En Go, no existen las enumeraciones (enums) nativas.
// En su lugar, se usa una combinación de constantes y `iota`.

// Se define un nuevo tipo llamado PizzaSize, que es un alias de int.
// Esto nos da seguridad de tipo.
type PizzaSize int

// `iota` es un generador de secuencias que se incrementa automáticamente.
// Cada constante en este bloque obtiene un valor secuencial, comenzando desde 0.
const (
	// Small es 0
	Small PizzaSize = iota
	// Medium es 1
	Medium
	// Large es 2
	Large
)

func Enums() {
	// Declaramos una variable para una pizza mediana, usando el "enum"
	myPizza := Medium

	// Usamos un switch para decidir qué acción tomar según el tamaño de la pizza.
	// Esto es más legible que usar números directamente (e.g., `case 1`).
	fmt.Println("Ordenando tu pizza...")

	switch myPizza {
	case Small:
		fmt.Println("Has ordenado una pizza pequeña.")
	case Medium:
		fmt.Println("Has ordenado una pizza mediana. ¡Excelente elección!")
	case Large:
		fmt.Println("Has ordenado una pizza grande. ¡Perfecta para compartir!")
	default:
		// Esta rama se ejecuta si el valor de myPizza no coincide con ninguno de los enums
		fmt.Println("Lo siento, ese tamaño de pizza no está disponible.")
	}

	// También podemos ver los valores numéricos detrás de los "enums"
	fmt.Printf("\nEl valor numérico de 'Small' es: %d\n", Small)
	fmt.Printf("El valor numérico de 'Medium' es: %d\n", Medium)
	fmt.Printf("El valor numérico de 'Large' es: %d\n", Large)
}
