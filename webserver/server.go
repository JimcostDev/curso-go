package webserver

import (
	"fmt"
	"net/http"
	"path/filepath"
)

//Un handler es una función que se encarga de procesar una petición HTTP y generar una respuesta para el cliente.
// w http.ResponseWriter: es una interfaz que se usa para responder al cliente HTTP (enviar datos).
//r *http.Request: es un puntero a una estructura que representa la petición HTTP recibida (lo que envía el cliente).

// handler para la ruta "/hola"
func holaHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "¡Hola desde tu primer servicio web en Go!")
}

// handler para la ruta "/saludo"
func saludoHandler(w http.ResponseWriter, r *http.Request) {
	nombre := r.URL.Query().Get("nombre")
	if nombre == "" {
		nombre = "visitante"
	}
	fmt.Fprintf(w, "Hola, %s 👋\n", nombre) // /saludo?nombre=JimcostDev
}

// handler para la ruta "/" que muestra un HTML
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	path := filepath.Join(".", "webserver", "index.html")
	http.ServeFile(w, r, path)
}

// Start inicia el servidor web
func Start() {
	// Rutas registradas
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hola", holaHandler)
	http.HandleFunc("/saludo", saludoHandler)

	fmt.Println("Servidor escuchando en http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("Error al iniciar el servidor: " + err.Error())
	}
}
