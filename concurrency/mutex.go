package concurrency

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func Mutex() {
	// Declaramos las variables compartidas
	var (
		contador int            // 🧮 Variable compartida que será protegida
		mu       sync.Mutex     // 🔐 Mutex para controlar acceso concurrente
		wg       sync.WaitGroup // 🚦 WaitGroup para esperar goroutines
	)

	// Información sobre núcleos disponibles
	fmt.Printf("💻 Núcleos disponibles: %d\n", runtime.NumCPU())

	// Lanzamos 5 goroutines para incrementar el contador
	for i := 0; i < 5; i++ {
		wg.Add(1) // 📌 Añadimos una tarea al contador de WaitGroup

		go func(id int) {
			defer wg.Done() // ✅ Marcamos tarea como completada al finalizar

			// 🛑 SECCIÓN CRÍTICA (comienza)
			mu.Lock() // 🔒 BLOQUEO - Solo UNA goroutine puede entrar aquí a la vez

			// 📝 Leemos el valor actual
			temp := contador
			fmt.Printf("🧵 Goroutine %d: valor actual = %d\n", id, temp)

			// Simulamos un pequeño retraso para demostrar la condición de carrera
			time.Sleep(100 * time.Millisecond)

			// ✏️ Modificamos el valor
			contador = temp + 1
			fmt.Printf("🆕 Goroutine %d: nuevo valor = %d\n", id, contador)

			mu.Unlock() // 🔓 DESBLOQUEO - Libera el acceso para otras goroutines
			// 🟢 FIN SECCIÓN CRÍTICA

			fmt.Printf("✅ Goroutine %d: incremento completado\n", id)
		}(i) // Pasamos el índice como parámetro
	}

	// Esperamos a que todas las goroutines terminen y mostramos el resultado final
	wg.Wait()
	fmt.Printf("\n🧮 CONTADOR FINAL: %d (esperado: 5)\n", contador)
	fmt.Println("💡 Conclusión: El mutex previene condiciones de carrera (ocurren cuando múltiples procesos o hilos intentan acceder y modificar un mismo recurso compartido al mismo tiempo).")
}
