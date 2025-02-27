package data_structure

// En Go, make() es una función incorporada que se utiliza para
// inicializar y asignar memoria para ciertos tipos de datos: slices, maps y channels.
import "fmt"

func Maps() {
	// Con make(): inicializa y asigna memoria
	goleadores := make(map[string]int)
	goleadores["Ronaldo"] = 30
	fmt.Println(goleadores)

	// Sin make():
	asistentes := map[string]int{}
	asistentes["James"] = 16
	fmt.Println(asistentes)

	// Estructuras de datos clave-valor
	equipos := map[string]string{
		"Barcelona":     "azulgrana",
		"Real Madrid":   "blanco",
		"Bayern Munich": "rojo",
	}
	// Acceder a un valor
	fmt.Println(equipos["Barcelona"])

	// Agregar un nuevo equipo
	equipos["Liverpool"] = "rojo"

	// verificar si un valor existe
	color, existe := equipos["Liverpool"]
	if existe {
		fmt.Println("Color de Liverpool: ", color)
	} else {
		fmt.Println("Liverpool no existe en el mapa")
	}

	// Recorrer el mapa
	fmt.Println("Equipos y sus colores:")
	for equipo, color := range equipos {
		fmt.Printf("El equipo %s tiene el color %s\n", equipo, color)
	}

	// Eliminar un valor
	delete(equipos, "Liverpool")

	// Imprimir el mapa actualizado
	fmt.Println("Mapa actualizado:", equipos)
}
