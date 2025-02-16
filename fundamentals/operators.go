package fundamentals

import (
	"fmt"
	"strings"
)

// operadores aritméticos
func arithmeticOperators() {
	fmt.Println(":: Operadores aritméticos ::")
	fmt.Println(10 + 5) // 15
	fmt.Println(10 - 5) // 5
	fmt.Println(10 * 5) // 50
	fmt.Println(10 / 5) // 2
	fmt.Println(10 % 5) // 0
	fmt.Print(strings.Repeat("-", 20))
}

// operadores de comparación
func comparisonOperators() {
	fmt.Println("")
	fmt.Println(":: Operadores de comparación ::")
	fmt.Println(10 == 5)  // false
	fmt.Println(10 == 10) // true
	fmt.Println(10 != 5)  // true
	fmt.Println(10 > 5)   // true
	fmt.Println(10 < 5)   // false
	fmt.Println(10 >= 5)  // true
	fmt.Println(10 <= 5)  // false
	fmt.Print(strings.Repeat("-", 20))
}

// operadores lógicos
func logicalOperators() {
	fmt.Println("")
	fmt.Println(":: Operadores lógicos ::")
	fmt.Println("AND:")
	fmt.Println(true && true)   // true
	fmt.Println(true && false)  // false
	fmt.Println(false && true)  // false
	fmt.Println(false && false) // false

	fmt.Println("")
	fmt.Println("OR:")
	fmt.Println(true || true)   // true
	fmt.Println(true || false)  // true
	fmt.Println(false || true)  // true
	fmt.Println(false || false) // false

	fmt.Println("")
	fmt.Println("NOT:")
	fmt.Println(!true)  // false
	fmt.Println(!false) // true
	fmt.Print(strings.Repeat("-", 20))
}

func Operators() {
	arithmeticOperators()
	comparisonOperators()
	logicalOperators()
}
