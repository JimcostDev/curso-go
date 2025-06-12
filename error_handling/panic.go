package error_handling

import "fmt"

func proteger() {
	if r := recover(); r != nil {
		fmt.Println("Se recuperó del pánico:", r)
	}
}

func lanzarPanic() {
	defer proteger()
	fmt.Println("Voy a causar un pánico...")
	panic("¡Algo salió mal!")
}

func Panic() {
	lanzarPanic()
	fmt.Println("El programa sigue funcionando.")
}
