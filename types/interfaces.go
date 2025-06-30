package types

import "github.com/JimcostDev/curso-go/types/models"

// Una interfaz define un conjunto de métodos sin implementación,
// estableciendo un contrato para los tipos que la cumplen.

// Conceptos clave:
// 1. Contrato: Un tipo cumple una interfaz si implementa sus métodos.
// 2. Desacoplamiento: Permite usar funciones genéricas sin depender de tipos concretos.
// 3. Implementación Implícita: No requiere declaración explícita de implementación.
// 4. Polimorfismo: Diferentes tipos pueden ser tratados de manera uniforme.
// 5. Composición: Una interfaz puede incluir otras para mayor flexibilidad.

// Definir la interfaz Luchador
type Luchador interface {
	Pelear() // contrato
	Entrenar()
}

// Función que usa la interfaz, permite ejecutar una pelea y un entrenamiento para cualquier tipo que implemente Luchador.
// - Polimorfismo: Llama a los métodos Pelear() y Entrenar() sin importar el tipo concreto.
// - Desacoplamiento: No depende de una implementación específica, solo de la interfaz Luchador.
func combate(l Luchador) {
	l.Pelear()
	l.Entrenar()
}

func Interfaces() {
	// Crear luchadores
	striker := models.Striker{
		Fighter: models.Fighter{Nombre: "Alex Pereira"},
		Strikes: 24,
	}
	grappler := models.Grappler{
		Fighter:   models.Fighter{Nombre: "Khamzat Chimaev"},
		Takedowns: 7,
	}
	allRounder := models.AllRounder{
		Fighter:  models.Fighter{Nombre: "Ilia Topuria"},
		Striker:  models.Striker{Strikes: 22},
		Grappler: models.Grappler{Takedowns: 2}}

	// Usar la interfaz en un combate
	// Como Striker, Grappler y AllRounder tienen un método Pelear(), automáticamente implementan Luchador.
	combate(striker)
	combate(grappler)
	combate(allRounder)
}
