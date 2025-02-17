// Función con retornos nombrados
package functions

import "fmt"

// Usamos retornos nombrados: area y perimetro
func calcularRectangulo(base, altura float64) (area float64, perimetro float64) {
	// Los valores area y perimetro ya están declarados por ser retornos nombrados
	// Podemos usarlos directamente como variables
	area = base * altura
	perimetro = 2 * (base + altura)

	// Al usar 'return' sin valores, automáticamente devuelve
	// los valores de las variables 'area' y 'perimetro'
	return
}

func NamedReturns() {
	fmt.Println("Función con retornos nombrados:")
	base := 5.0
	altura := 3.0

	area, perimetro := calcularRectangulo(base, altura)
	fmt.Printf("Rectángulo de %.1f x %.1f:\n", base, altura)
	fmt.Printf("Área: %.2f\n", area)
	fmt.Printf("Perímetro: %.2f\n", perimetro)

	// Otro ejemplo usando diferentes medidas
	a2, p2 := calcularRectangulo(10.0, 4.0)
	fmt.Printf("\nRectángulo de 10.0 x 4.0:\n")
	fmt.Printf("Área: %.2f\n", a2)
	fmt.Printf("Perímetro: %.2f\n", p2)
}
