package unittests

// La función Sumar suma los elementos de una lista de enteros
func SumarEnteros(lista []int) int {
	total := 0
	for _, num := range lista {
		total += num
	}
	return total
}
