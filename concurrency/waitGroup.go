package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func WaitGroup() {
	// Creamos un WaitGroup para coordinar nuestras goroutines
	var wg sync.WaitGroup

	fmt.Println("🚀 Iniciando programa...")
	fmt.Println("-----------------------")

	// Vamos a lanzar 2 goroutines
	wg.Add(2)

	// Goroutine 1: Saludo rápido
	go func() {
		defer wg.Done() // 📍 Marcamos finalización al terminar

		// Simulamos un pequeño procesamiento
		time.Sleep(300 * time.Millisecond)
		fmt.Println("👋 ¡Hola desde la goroutine #1!")
	}()

	// Goroutine 2: Saludo con más trabajo
	go func() {
		defer wg.Done() // 📍 Marcamos finalización al terminar

		// Simulamos un procesamiento más largo
		time.Sleep(600 * time.Millisecond)
		fmt.Println("🌟 ¡Hola desde la goroutine #2! (tuve más trabajo)")
	}()

	// Mensaje desde el hilo principal
	fmt.Println("🖥️  Mensaje desde main (hilo principal)")
	fmt.Println("💤 Main: voy a esperar a que las goroutines terminen...")
	fmt.Println("-----------------------")

	// Esperamos a que todas las goroutines finalicen
	wg.Wait()

	fmt.Println("-----------------------")
	fmt.Println("🎉 ¡Todas las goroutines han terminado!")
	fmt.Println("🏁 Programa finalizado.")
}
