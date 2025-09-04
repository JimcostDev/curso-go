package main

import (
	"github.com/JimcostDev/curso-go/ai"
	"github.com/JimcostDev/curso-go/concurrency"
	"github.com/JimcostDev/curso-go/exercises"
	"github.com/JimcostDev/curso-go/fundamentals"
	"github.com/JimcostDev/curso-go/utils"
)

func main() {
	// aqui probamos nuestro código
	fundamentals.Hello()
	ai.Gemini()
	utils.TitleCaseExample("hola mundo")
	utils.Number(5)
	exercises.PrintPy()
	concurrency.Select()
}
