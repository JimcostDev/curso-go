package error_handling

import (
	"fmt"
)

func saludar() {
	defer fmt.Println("Adiós")
	fmt.Println("Hola")
}

func Defer() {
	saludar()
}
