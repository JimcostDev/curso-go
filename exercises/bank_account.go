package exercises

import (
	"fmt"
)

// Definimos la estructura
type CuentaBancaria struct {
	saldo float64
}

// Método para depositar
func (c *CuentaBancaria) Depositar(monto float64) {
	if monto > 0 {
		c.saldo += monto
		fmt.Printf("Has depositado %.2f. Nuevo saldo: %.2f\n", monto, c.saldo)
	} else {
		fmt.Println("El monto a depositar debe ser mayor que 0.")
	}
}

// Método para retirar
func (c *CuentaBancaria) Retirar(monto float64) {
	if monto <= 0 {
		fmt.Println("El monto a retirar debe ser mayor que 0.")
	} else if monto > c.saldo {
		fmt.Println("Fondos insuficientes.")
	} else {
		c.saldo -= monto
		fmt.Printf("Has retirado %.2f. Nuevo saldo: %.2f\n", monto, c.saldo)
	}
}

// Método para mostrar saldo
func (c *CuentaBancaria) MostrarSaldo() {
	fmt.Printf("Saldo actual: %.2f\n", c.saldo)
}

func BankAccount() {
	// Crear cuenta con 100 iniciales
	nequi := CuentaBancaria{saldo: 1000}

	// Usar los métodos
	nequi.MostrarSaldo()
	nequi.Depositar(500)
	nequi.Retirar(300)
	nequi.Retirar(200)
	nequi.MostrarSaldo()
}
