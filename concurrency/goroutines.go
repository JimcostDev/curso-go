package concurrency

import (
	"fmt"
	"time"
)

func Goroutines() {
	// Lanzamos dos goroutines (ejecución concurrente)
	go tarea("🌞 Desayunar", 2)   // 2 segundos
	go tarea("📚 Estudiar Go", 4) // 4 segundos

	// Esperamos para que no termine el programa
	time.Sleep(5 * time.Second)
	fmt.Println("✅ Todas las tareas completadas")
}

func tarea(nombre string, segundos int) {
	fmt.Printf("▶️ Iniciando: %s\n", nombre)
	time.Sleep(time.Duration(segundos) * time.Second)
	fmt.Printf("🏁 Terminado: %s\n", nombre)
}
