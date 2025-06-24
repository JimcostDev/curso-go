/*
	El Índice de Masa Corporal (IMC) es una medida que relaciona el peso y
	la estatura de una persona para estimar si tiene un peso adecuado.
	Se calcula con la fórmula:
	IMC = peso (kg) / estatura^2 (m)

	Rango de IMC (clasificación según la OMS):
	- Bajo peso: IMC < 18.5
	- Normal: IMC 18.5–24.9
	- Sobrepeso: IMC 25–29.9
	- Obesidad: IMC >= 30 (Obesidad tipo I: 30–34.9, Obesidad tipo II: 35–39.9, Obesidad tipo III: >= 40)
*/

package exercises

import (
	"fmt"
	"strconv"
)

// CalcularIMC calcula el Índice de Masa Corporal (IMC) de una persona
func calcularIMC(peso, estatura float64) float64 {
	return peso / (estatura * estatura)
}

// ClasificarIMC clasifica el IMC según la OMS
func clasificarIMC(imc float64) string {
	switch {
	case imc < 18.5:
		return "Bajo peso"
	case imc < 25:
		return "Normal"
	case imc < 30:
		return "Sobrepeso"
	case imc < 35:
		return "Obesidad tipo I"
	case imc < 40:
		return "Obesidad tipo II"
	default:
		return "Obesidad tipo III"
	}
}

// Calcular peso maximo dentro del rango de IMC normal
func calcularPesoMaximo(estatura float64) float64 {
	return 24.9 * estatura * estatura
}

// Calcular peso minimo dentro del rango de IMC normal
func calcularPesoMinimo(estatura float64) float64 {
	return 18.5 * estatura * estatura
}

// Mostar peso maximo y minimo dentro del rango de IMC normal
func mostrarPesoNormal(estatura float64) {
	fmt.Printf("Peso normal: %.2f - %.2f kg\n", calcularPesoMinimo(estatura), calcularPesoMaximo(estatura))
}

// Mostrar resultados del IMC
func mostrarIMC(imc float64, clasificacion string) {
	fmt.Printf("Tu IMC es: %.2f\n", imc)
	fmt.Println("Clasificación:", clasificacion)
}

func IMC() {
	var pesoInput, alturaInput string

	// Leer el peso
	fmt.Print("Introduce tu peso en kg (por ejemplo, 66.1): ")
	fmt.Scanln(&pesoInput)

	// Convertir peso a float64
	peso, err := strconv.ParseFloat(pesoInput, 64)
	if err != nil {
		fmt.Println("Error: Debes introducir un peso válido (número decimal).")
		return
	}

	// Leer la altura
	fmt.Print("Introduce tu altura en metros (por ejemplo, 1.84): ")
	fmt.Scanln(&alturaInput)

	// Convertir altura a float64
	estatura, err := strconv.ParseFloat(alturaInput, 64)
	if err != nil {
		fmt.Println("Error: Debes introducir una altura válida (número decimal).")
		return
	}

	// Calcular IMC
	imc := calcularIMC(peso, estatura)

	// Clasificar IMC
	clasificacion := clasificarIMC(imc)

	// Mostrar resultados
	mostrarIMC(imc, clasificacion)

	// Mostrar peso normal
	mostrarPesoNormal(estatura)
}
