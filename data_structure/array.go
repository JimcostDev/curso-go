package data_structure

import "fmt"

func Array() {
	// tradicional
	// var teams [3]string
	// teams[0] = "Cali"
	// teams[1] = "Liverpool"
	// teams[2] = "Milan"

	teams := [...]string{"Cali", "Liverpool", "Milan", "Barcelona"}
	fmt.Println(teams)

	// recorrer array
	for _, team := range teams {
		fmt.Println(team)
	}

	fmt.Printf("El tipo de dato es: %T\n", teams)
}
