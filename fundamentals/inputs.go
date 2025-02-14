package fundamentals

// Para valores simples → Scan : Lee valores ignorando espacios y saltos de línea
// Para valores en una línea → Scanln: Lee solo hasta el salto de línea
// Para formatos específicos → Scanf: Permite especificar un formato exacto
// Para texto con espacios → bufio.NewReader: Lee líneas completas incluyendo espacios

import (
	"bufio"
	"fmt"
	"os"
)

func Inputs() {
	// 1. Scan - Lectura básica
	var edad int
	fmt.Print("Ingresa tu edad: ")
	fmt.Scan(&edad)
	fmt.Printf("Edad leída con Scan: %d\n\n", edad)

	// 2. Scanln - Dos valores en una línea
	reader := bufio.NewReader(os.Stdin) // Crear un lector
	reader.ReadString('\n')             // limpiar el buffer
	var dia, mes int
	fmt.Print("Ingresa día y mes (ejemplo: 15 3): ")
	fmt.Scanln(&dia, &mes)
	fmt.Printf("Fecha leída con Scanln: %d/%d\n\n", dia, mes)

	// 3. Scanf - Formato específico
	var horas, minutos int
	fmt.Print("Ingresa la hora (formato HH:MM): ")
	fmt.Scanf("%d:%d", &horas, &minutos)
	fmt.Printf("Hora leída con Scanf: %02d:%02d\n\n", horas, minutos)

	// 4. bufio.NewReader - Texto con espacios
	reader.ReadString('\n') // Limpiar el buffer
	fmt.Print("Ingresa otra frase: ")
	frase, _ := reader.ReadString('\n') // Leer hasta el salto de línea
	//frase = strings.TrimSpace(frase)    // Limpiar espacios al inicio y final
	fmt.Printf("[%s]\n", frase)
	fmt.Printf("Longitud: %d\n", len(frase))
}
