package error_handling

import (
	"fmt"
)

// En Go, un error personalizado se crea implementando el método `Error()`
// de la interfaz `error`.

// CoffeeError representa un error específico al ordenar café
type CoffeeError struct {
	Code    int    // Un código para identificar el tipo de error
	Message string // Descripción del error
}

// El método `Error()` convierte CoffeeError en un error válido de Go
func (e CoffeeError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

// orderCoffee intenta preparar un café, pero puede fallar
func orderCoffee(size string) error {
	// Simulamos algunos casos de error
	if size == "" {
		return CoffeeError{Code: 400, Message: "Debes elegir un tamaño de café."}
	}
	if size == "extra-grande" {
		return CoffeeError{Code: 404, Message: "Ese tamaño de café no existe."}
	}
	if size == "expreso" {
		return CoffeeError{Code: 503, Message: "La máquina de expreso está en mantenimiento."}
	}

	// Si no hay error, todo sale bien
	fmt.Printf("☕ Preparando tu café tamaño '%s'...\n", size)
	return nil
}

func Coffee() {
	// Intentamos ordenar varios cafés para ver los diferentes resultados
	sizes := []string{"", "extra-grande", "expreso", "mediano"}

	for _, s := range sizes {
		err := orderCoffee(s)
		if err != nil {
			// Aquí mostramos el error personalizado
			fmt.Println("❌ Ocurrió un problema:", err)
		} else {
			fmt.Println("✅ ¡Tu café está listo!")
		}
		fmt.Println("-----------------------------")
	}
}
