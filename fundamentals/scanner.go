package fundamentals

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Input() {
	// captura nombre y edad
	var nombre string
	var edad int
	fmt.Print("Ingresa tu nombre y edad: ")
	fmt.Scan(&nombre, &edad)
	fmt.Printf("Nombre: %s, Edad: %d\n", nombre, edad)

	// Limpiar el buffer de entrada
	bufio.NewReader(os.Stdin).ReadString('\n')

	// captura frase
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Escribe tu frase favorita: ")
	scanner.Scan()
	frase := strings.TrimSpace(scanner.Text()) // elimina espacios en blanco
	fmt.Println("Tu frase es:", frase)
}
