package models

import "fmt"

// Fighter contiene los campos comunes para todos los luchadores.
type Fighter struct {
	Nombre string
	// Aqui podemos agregar otros campos comunes para los luchadores
}

// Método Entrenar para Fighter.
func (f Fighter) Entrenar() {
	fmt.Printf("%s está entrenando!\n", f.Nombre)
}

// Striker: Especialista en pelea de pie. Utiliza técnicas de boxeo, muay thai, kickboxing, karate, etc. Busca el nocaut o dominar desde la distancia.
type Striker struct {
	Fighter // Embebe Fighter, heredando el campo Nombre.
	Strikes int
}

// Método Pelear para Striker.
func (s Striker) Pelear() {
	fmt.Printf("%s lanza %d strikes!\n", s.Nombre, s.Strikes)
}

// Grappler: Especialista en lucha cuerpo a cuerpo. Se enfoca en derribos, control en el suelo, sumisiones. Proviene del wrestling, jiu-jitsu, judo, sambo, etc.
type Grappler struct {
	Fighter
	Takedowns int
}

// Método Pelear para Grappler.
func (g Grappler) Pelear() {
	fmt.Printf("%s realiza %d derribos!\n", g.Nombre, g.Takedowns)
}

// AllRounder: Versátil. Tiene habilidades sólidas tanto de striking como de grappling. Puede adaptarse a distintos rivales y estilos.
// Embebe tanto a Striker como a Grappler.
type AllRounder struct {
	Fighter
	Striker
	Grappler
}

// Método Pelear para AllRounder.
func (a AllRounder) Pelear() {
	fmt.Printf("%s combina el striking y el grappling: %d strikes y %d derribos!\n",
		a.Nombre,
		a.Strikes,
		a.Takedowns)
}
