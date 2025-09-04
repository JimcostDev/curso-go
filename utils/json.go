package utils

import (
	"encoding/json"
	"fmt"
)

// Creamos una estructura Go con etiquetas JSON.
// Las etiquetas controlan cómo se mapean los campos a JSON.
// `json:"nombre"` especifica que el campo se llamará "nombre" en el JSON.
// `json:"-"` omite el campo del JSON.
// `json:"edad,omitempty"` omite el campo si su valor es el valor cero (ej. 0 para int, "" para string).
type Persona struct {
	Nombre string `json:"nombre"`
	Edad   int    `json:"edad"`
	Ciudad string `json:"ciudad,omitempty"`
	ID     string `json:"-"`
}

func Encode() {
	// 1. Crear una instancia de la estructura.
	p := Persona{
		Nombre: "Ronaldo",
		Edad:   26,
		Ciudad: "Cali",
		ID:     "12345",
	}

	// 2. Codificar la estructura a JSON.
	// La función Marshal devuelve un slice de bytes y un error.
	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error al codificar:", err)
		return
	}

	// 3. Imprimir el JSON resultante (como string).
	fmt.Println("JSON codificado:", string(jsonData))
}

func Decode() {
	// 1. JSON de ejemplo (como slice de bytes).
	jsonData := []byte(`{"nombre":"Ronaldo","edad":26,"ciudad":"Cali"}`)

	// 2. Crear una variable para almacenar la estructura decodificada.
	// Se pasa el slice de bytes y un puntero a la estructura.
	var p Persona
	err := json.Unmarshal(jsonData, &p)
	if err != nil {
		fmt.Println("Error al decodificar:", err)
		return
	}

	// 3. Imprimir la estructura decodificada.
	fmt.Println("Estructura decodificada:")
	fmt.Println("Nombre:", p.Nombre)
	fmt.Println("Edad:", p.Edad)
	fmt.Println("Ciudad:", p.Ciudad)
}
