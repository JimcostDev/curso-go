package data_structure

/*
1. Los arrays trabajan por valor, lo que significa que al asignar o pasar un array a una función,
se crea una copia completa del array. Los cambios realizados en la copia no afectan al array original.

2. Los slices trabajan por referencia, lo que significa que cuando asignas o pasas un slice a
una función, cualquier cambio en el slice afecta al array subyacente, ya que el slice solo es una vista del array.
*/
import "fmt"

func modificarArray(arr [3]string) {
	arr[0] = "Modificado en función"
	fmt.Println("Array dentro de la función:", arr)
}

func modificarSlice(slc []string) {
	slc[0] = "Modificado en función"
	fmt.Println("Slice dentro de la función:", slc)
}

func Diff() {

	// Declaración de un array
	array := [3]string{"Original 1", "Original 2", "Original 3"}

	// Declaración de un slice
	slice := []string{"Original 1", "Original 2", "Original 3"}

	// Pasar array y slice a las funciones
	modificarArray(array)
	modificarSlice(slice)

	// Imprimir resultados
	fmt.Println("Array después de modificar:", array)
	fmt.Println("Slice después de modificar:", slice)

}
