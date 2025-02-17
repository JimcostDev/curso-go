package flow_control

import (
	"fmt"
)

// los ciclos en go son muy parecidos a los de otros lenguajes
// nos permiten repetir un bloque de código mientras se cumpla una condición

func Iterators() {

	// tradicional
	for i := 0; i < 5; i++ {
		fmt.Println("Iteración:", i)
	}

	// equivalente a un while
	x := 0
	for x < 5 {
		fmt.Println("Valor de x:", x)
		x++
	}

	// infinito
	// for {
	// 	fmt.Println(x)
	// 	x++
	// }

	// range: nos permite iterar sobre un slice, array, string, map o channel
	frutas := []string{"piña", "mango", "mandarina"}
	for idx, fruta := range frutas {
		fmt.Printf("Índice: %d, Valor: %s\n", idx, fruta)
	}

	// mapa o diccionario
	colores := map[string]string{
		"rojo":  "#FF0000",
		"verde": "#00FF00",
		"azul":  "#0000FF",
	}
	for clave, valor := range colores {
		fmt.Printf("Color: %s, Código: %s\n", clave, valor)
	}

	// continue
	for i := 0; i < 10; i++ {
		if i%2 == 0 { // Saltar números pares
			continue
		}
		fmt.Println("Número impar:", i)
	}

	// break
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("¡Llegamos a 5!")
			break
		}
		fmt.Println("Iteración:", i)
	}

}
