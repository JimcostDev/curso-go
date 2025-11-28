package exercises

import (
	"fmt"
	"math"
	"strconv"
)

// Constantes para los límites del IMC, siguiendo la clasificación de la OMS.
const (
	BajoPesoLimite  = 18.5
	NormalLimite    = 25.0
	SobrepesoLimite = 30.0
	Obesidad1Limite = 35.0
	Obesidad2Limite = 40.0
	// Usamos 24.9 y 18.5 para el cálculo de peso Normal
	IMCMaxNormal = 24.9
	IMCMinNormal = 18.5
)

// CalcularIMC calcula el Índice de Masa Corporal (IMC) de una persona.
// Retorna 0 si la estatura es cero para evitar división por cero.
func calcularIMC_r(peso, estatura float64) float64 {
	if estatura <= 0 {
		return 0
	}
	return peso / math.Pow(estatura, 2)
}

// ClasificarIMC clasifica el IMC según la OMS usando las constantes definidas.
func clasificarIMC_r(imc float64) string {
	switch {
	case imc < BajoPesoLimite:
		return "Bajo peso"
	case imc < NormalLimite: // Menor que 25.0 (Rango Normal: 18.5–24.9)
		return "Normal"
	case imc < SobrepesoLimite: // Menor que 30.0 (Rango Sobrepeso: 25.0–29.9)
		return "Sobrepeso"
	case imc < Obesidad1Limite: // Menor que 35.0 (Obesidad tipo I: 30.0–34.9)
		return "Obesidad tipo I"
	case imc < Obesidad2Limite: // Menor que 40.0 (Obesidad tipo II: 35.0–39.9)
		return "Obesidad tipo II"
	default: // IMC >= 40.0
		return "Obesidad tipo III"
	}
}

// calcularPesoLimite calcula el peso (kg) para un IMC límite dado a una estatura (m).
func calcularPesoLimite(estatura float64, imcLimite float64) float64 {
	peso_limite := imcLimite * math.Pow(estatura, 2)
	return peso_limite
}

// Mostrar resultados del IMC
func mostrarIMC_r(imc float64, clasificacion string) {
	fmt.Printf("\n--- Resultados ---\n")
	fmt.Printf("Tu IMC es: **%.2f**\n", imc)
	fmt.Println("Clasificación:", clasificacion)
}

// Mostrar peso maximo y minimo dentro del rango de IMC normal
func mostrarPesoNormal_r(estatura float64) {
	min := calcularPesoLimite(estatura, IMCMinNormal)
	max := calcularPesoLimite(estatura, IMCMaxNormal)
	fmt.Printf("\nPara tu altura (%.2f m):\n", estatura)
	fmt.Printf("⚖️  Rango de peso **Normal** (IMC 18.5-24.9) es: **%.2f - %.2f kg**\n", min, max)
}

// solicitarEntrada lee una entrada de la consola y la convierte a float64.
// Retorna el valor y un posible error.
func solicitarEntrada(mensaje string) (float64, error) {
	var input string
	fmt.Print(mensaje)
	fmt.Scanln(&input)

	valor, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, fmt.Errorf("entrada no válida. Por favor, introduce un número decimal, por ejemplo: 66.1 o 1.84. Error: %w", err)
	}
	return valor, nil
}

// IMC es la función principal que coordina la lógica del programa.
func IMC_r() {
	// 1. Leer y validar el peso
	peso, err := solicitarEntrada("Introduce tu peso en kg (Ej. 66.1): ")
	if err != nil {
		fmt.Println("❌ Error en el peso:", err)
		return
	}
	if peso <= 0 {
		fmt.Println("❌ Error: El peso debe ser un valor positivo.")
		return
	}

	// 2. Leer y validar la altura
	estatura, err := solicitarEntrada("Introduce tu altura en metros (Ej. 1.84): ")
	if err != nil {
		fmt.Println("❌ Error en la altura:", err)
		return
	}
	if estatura <= 0 {
		fmt.Println("❌ Error: La altura debe ser un valor positivo.")
		return
	}

	// 3. Calcular, clasificar y mostrar resultados
	imc := calcularIMC_r(peso, estatura)
	clasificacion := clasificarIMC_r(imc)
	mostrarIMC_r(imc, clasificacion)
	mostrarPesoNormal_r(estatura)
}
