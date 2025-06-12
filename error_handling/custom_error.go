package error_handling

import (
	"fmt"
)

// Definimos un tipo que implementa la interfaz error
type ErrorPersonalizado struct {
	Mensaje string
	Codigo  int
}

// Implementamos el método Error() para que cumpla con la interfaz error
func (e ErrorPersonalizado) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Codigo, e.Mensaje)
}

func hacerAlgo(peligroso bool) error {
	if peligroso {
		return ErrorPersonalizado{
			Mensaje: "ocurrió un problema grave",
			Codigo:  500,
		}
	}
	return nil
}

func CustomError() {
	err := hacerAlgo(true) // Cambiamos a true para simular un error
	// err := hacerAlgo(false) // Cambiamos a false para simular éxito

	if err != nil {
		fmt.Println("Ocurrió un error:", err)
		// Podemos usar type assertion si queremos acceder al error como struct
		if e, ok := err.(ErrorPersonalizado); ok {
			fmt.Println("Código del error:", e.Codigo)
		}
		return
	}

	fmt.Println("Todo salió bien.")
}
