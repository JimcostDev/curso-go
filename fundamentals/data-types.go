package fundamentals

// Tipos de datos básicos en Go:
// Números enteros (int, int8, int16, int32, int64)
// Números decimales (float32, float64)
// Booleanos (bool)
// Cadenas de texto (string)
// Runa (caracteres) (rune)
// Byte (equivalente a uint8) (byte)
// tipos de datos en Go
import (
	"fmt"
)

func DataTypes() {
	// Enteros
	var entero int = 42
	var pequeño int8 = 127 // -128 a 127

	// Decimales
	var decimal float32 = 3.14
	var grande float64 = 2.7182818284

	// Booleano
	var verdadero bool = true

	// String
	var texto string = "Hola, Go!"

	// Rune y Byte
	var caracter rune = '@'
	var letra byte = '@'

	fmt.Println("Entero:", entero)
	fmt.Println("Pequeño:", pequeño)
	fmt.Println("Decimal:", decimal)
	fmt.Println("Grande:", grande)
	fmt.Println("Booleano:", verdadero)
	fmt.Println("Texto:", texto)
	fmt.Println("Caracter:", caracter)
	fmt.Println("Letra (byte):", letra)
	Hello() // llamamos a la función Hello

	// inferencia de tipos
	edad := 25             // Go infiere que es int
	nombre := "JimcostDev" // Go infiere que es string
	fmt.Println("Edad:", edad)
	fmt.Println("Nombre:", nombre)

	// Si declaramos una variable sin asignarle un valor, Go le asigna su zero value:
	var sinValor int    // 0
	var sinTexto string // ""
	var sinBool bool    // false
	fmt.Println("Sin valor:", sinValor)
	fmt.Println("Sin texto:", sinTexto)
	fmt.Println("Sin booleano:", sinBool)

	//  Go no hay conversión automática, se debe hacer manualmente:
	var x int = 42
	var y float64 = float64(x) // Convertimos int a float64
	fmt.Println(y)

	// constantes
	const pi float64 = 3.1416
	const gravedad float64 = 9.8
	fmt.Println("Pi:", pi)
	fmt.Println("Gravedad:", gravedad)
}
