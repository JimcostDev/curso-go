package fundamentals

// Los apuntadores (o punteros) permiten almacenar la dirección de memoria de una variable.
// En Go, los apuntadores se representan con el símbolo * seguido del tipo de dato al que apunta.
// Para obtener la dirección de memoria de una variable, se utiliza el símbolo & seguido del nombre de la variable.
// Solo pueden apuntar a un tipo en concreto, no se pueden mezclar tipos de datos.
import "fmt"

func Pointers() {
	fmt.Println("Punteros")
	var puntero *int
	numero := 10
	puntero = &numero // direccion de memoria

	fmt.Printf("El valor de la variable es %v\n", numero)
	fmt.Printf("La dirección de memoria de la variable es %v\n", &numero)

	fmt.Printf("El valor de la variable a la que apunta el puntero es %v\n", *puntero)
	fmt.Printf("La dirección de memoria de la variable a la que apunta el puntero es %v\n", puntero)

	// cambiar el valor de la variable
	*puntero = 20
	fmt.Printf("El nuevo valor de la variable es %v\n", numero)
}
