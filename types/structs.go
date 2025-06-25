package types

import "fmt"

// Un struct es una colección de campos. Cada campo tiene un nombre y un tipo.
// Es decir puedo grupar distintas propiedades bajo un mismo tipo,
// como si fuera una “plantilla” u “objeto” en otros lenguajes.
// Además, puedes asociar funciones (llamadas métodos) a estos
// tipos, para que operen sobre sus datos internos.

// Un método es una función con un receptor (receiver).
// El receptor indica sobre qué tipo de dato va a operar esa función.

// definir un struct
type ArtistaMarcial struct {
	Nombre    string
	Peso      float64
	Altura    float64
	Victorias int
	Derrotas  int
}

// Método para presentar al peleador
// Receptor por valor (a ArtistaMarcial). Esto significa que se trabaja con una copia del objeto original.
func (a ArtistaMarcial) Presentarse() {
	fmt.Printf("Soy %s, peso %.1f kg y mi récord es %d - %d\n", a.Nombre, a.Peso, a.Victorias, a.Derrotas)
}

// Método para registrar una victoria
// Receptor por puntero (a *ArtistaMarcial). Esto significa que el método recibe una referencia al objeto original, y puede modificar su estado interno
func (a *ArtistaMarcial) Victoria() {
	a.Victorias++
	fmt.Printf("¡%s ha ganado una pelea! Nuevo récord: %d - %d\n", a.Nombre, a.Victorias, a.Derrotas)
}

// Método para registrar una derrota
func (a *ArtistaMarcial) Derrota() {
	a.Derrotas++
	fmt.Printf("¡%s ha perdido una pelea! Nuevo récord: %d - %d\n", a.Nombre, a.Victorias, a.Derrotas)
}

// Se usa un puntero en el segundo método (*ArtistaMarcial) porque se necesita modificar el objeto original.
// En el primero no se usa porque solo se lee, y una copia es suficiente.

func Structs() {
	// crear un luchador
	luchador := ArtistaMarcial{
		Nombre:    "Ilia Topuria",
		Peso:      66,
		Altura:    1.70,
		Victorias: 16,
		Derrotas:  0,
	}

	luchador.Presentarse()
	luchador.Victoria()
}
