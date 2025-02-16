package flow_control

import (
	"fmt"
	"math/rand"
)

// bloques condicionales
func conditionals() {
	if valor := rand.Intn(10); valor == 0 {
		fmt.Println("El número", valor, "es cero")
	} else if valor%2 == 0 {
		fmt.Println("El número", valor, "es par")
	} else {
		fmt.Println("El número", valor, "es impar")
	}
	fmt.Println("Fin del programa")
}

// casos o switch
func switchCases() {
	dia := "domingo"

	switch dia {
	case "lunes":
		fmt.Println("Hoy es lunes, el primer día de la semana")
	case "martes", "miércoles", "jueves":
		fmt.Println("Mitad de de la semana")
	case "viernes":
		fmt.Println("Viernes, fin de la semana laboral")
	case "sábado", "domingo":
		fmt.Println("Fin de semana")
	default:
		fmt.Println("Día no válido")
	}

	letra := "a"

	switch letra {
	case "a", "e", "i", "o", "u", "A", "E", "I", "O", "U":
		fmt.Println("Es una vocal")
	default:
		fmt.Println("Es cualquier otro caracter  menos una vocal")
	}
}

func Selection() {
	//conditionals()
	switchCases()
}
