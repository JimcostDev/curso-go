package fundamentals

import "fmt"

type Kage string

const (
	Hokage     Kage = "Hokage"
	Kazekage   Kage = "Kazekage"
	Mizukage   Kage = "Mizukage"
	Raikage    Kage = "Raikage"
	Tsuchikage Kage = "Tsuchikage"
)

// tipos de datos personalizados
func CustomTypes() {
	Tsunade := Hokage
	fmt.Println("Tsunade es la: ", Tsunade)
}
