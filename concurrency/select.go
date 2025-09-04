package concurrency

import (
	"fmt"
	"time"
)

// Función principal que orquesta la preparación de las arepas.
func Select() {
	// Creamos canales para cada tarea, como bandejas de notificación.
	canalMasa := make(chan string)
	canalSarten := make(chan string)
	canalQueso := make(chan string)

	// Lanzamos las goroutines para cada tarea.
	// Cada una tiene un tiempo de preparación diferente.
	go preparar(canalMasa, "Masa lista para amasar", 3*time.Second)
	go preparar(canalSarten, "Sartén caliente", 1*time.Second)
	go preparar(canalQueso, "Queso rallado", 2*time.Second)

	fmt.Println("Comenzando a preparar las arepas...")
	fmt.Println("------------------------------------")

	// Usamos un bucle para esperar los 3 mensajes que necesitamos.
	for i := 0; i < 3; i++ {
		// El bloque `select` espera por el primer canal que reciba un mensaje.
		select {
		case mensaje := <-canalSarten:
			fmt.Printf("🔥 Recibido: %s. ¡A empezar a cocinar!\n", mensaje)
		case mensaje := <-canalMasa:
			fmt.Printf("🥣 Recibido: %s. ¡Ahora sí a darle forma!\n", mensaje)
		case mensaje := <-canalQueso:
			fmt.Printf("🧀 Recibido: %s. ¡Perfecto para el relleno!\n", mensaje)
		}
	}

	fmt.Println("------------------------------------")
	fmt.Println("¡Todo está listo! Hora de armar y disfrutar las arepas. 🎉")
}

// preparar simula una tarea que toma un tiempo y envía un mensaje a un canal cuando está lista.
func preparar(ch chan string, mensaje string, t time.Duration) {
	time.Sleep(t) // Simula el tiempo que toma la tarea.
	ch <- mensaje // Envía la notificación al canal.
}
