package utils

import (
	"fmt"

	// Paquetes de módulos externos (jimcostdev-utils)
	"github.com/JimcostDev/jimcostdev-utils/random"
	"github.com/JimcostDev/jimcostdev-utils/strutils"
)

func TitleCaseExample(input string) {
	result := strutils.ToTitleCase(input)
	fmt.Println(result)
}

func Number(length int) {
	result := random.UniqueDigits(length)
	fmt.Printf("Número aleatorio con %d dígitos únicos: %s\n", length, result)
}
