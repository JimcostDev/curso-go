package concurrency

import (
	"fmt"
	"time"
)

// enviarMensaje simula una tarea que se ejecuta de forma concurrente.
// Toma un canal de tipo 'string' como argumento para enviar un mensaje.
func enviarMensaje(canal chan string) {
	// Simula una tarea que toma tiempo, como una llamada a una API o un cálculo largo.
	fmt.Println("👉 Goroutine: Iniciando tarea y preparando el mensaje...")
	time.Sleep(3 * time.Second)

	// Aquí es donde ocurre la comunicación. El canal se bloquea hasta que
	// el programa principal esté listo para recibir este mensaje.
	fmt.Println("👉 Goroutine: ¡Mensaje listo! Enviando al canal.")
	canal <- "¡Hola desde la Goroutine! 👋"
	fmt.Println("👉 Goroutine: Mensaje enviado. La tarea ha finalizado.")
}

// Channels es la función principal que orquesta la comunicación.
func Channels() {
	// Paso 1: Crear un canal. Es como instalar la "tubería" de comunicación.
	// La capacidad del canal por defecto es 0, lo que significa que
	// la goroutine que envía y la que recibe deben estar listas
	// al mismo tiempo (comunicación síncrona).
	mensajes := make(chan string)

	// Paso 2: Lanzar la goroutine que se ejecutará en paralelo.
	// Le pasamos el canal que acabamos de crear para que pueda enviar su mensaje.
	go enviarMensaje(mensajes)

	// Paso 3: Esperar a recibir el mensaje del canal.
	// Esta línea de código es "bloqueante". El programa principal se detiene aquí
	// y espera hasta que haya un mensaje disponible en el canal.
	fmt.Println("--------------------------------------------------")
	fmt.Println("Main: Esperando a que la goroutine envíe su mensaje...")
	
	// '<-mensajes' es el operador para recibir un valor del canal.
	mensajeRecibido := <-mensajes

	fmt.Println("--------------------------------------------------")
	// Una vez que el mensaje se recibe, el programa continúa.
	fmt.Println("Main: ¡Mensaje recibido! 🎉")
	fmt.Println("Main: Contenido del mensaje:", mensajeRecibido)
}