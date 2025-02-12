package fundamentals

// Los generics en Go (introducidos en Go 1.18) permiten definir
// funciones y estructuras de datos que pueden trabajar con diferentes
//  tipos de datos sin necesidad de escribir código duplicado.
import "fmt"

func sumar[T int | float64](a T, b T) T {
	return a + b
}

func Generics() {
	rEnteros := sumar(5, 7)
	rFlotantes := sumar(5.5, 7.7)
	fmt.Printf("Suma de enteros: %d\n", rEnteros)
	fmt.Printf("Suma de flotantes: %f\n", rFlotantes)
}
