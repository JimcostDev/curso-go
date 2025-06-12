package error_handling

import (
	"errors"
	"fmt"
)

func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("no se puede dividir entre cero")
		// se retorna 0 por que es el valor por defecto de float64
	}
	return a / b, nil
}

func ShowError() {
	resultado, err := dividir(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Resultado:", resultado)
}
