package utils

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Bcrypt() {
	fmt.Println("--- SIMULACIÓN DEL REGISTRO DE USUARIO ---")

	// 1. Se define la contraseña del usuario.
	password := []byte("cambiar123")

	// 2. Se genera el hash de la contraseña.
	// bcrypt.GenerateFromPassword genera un hash que ya incluye la sal.
	// El segundo argumento es el 'cost', que determina la fuerza del hash.
	// bcrypt.DefaultCost es un valor por defecto seguro.
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	// Simulamos que `hashedPassword` es lo que se guarda en la base de datos.
	fmt.Printf("Contraseña en texto plano: %s\n", password)
	fmt.Printf("Hash almacenado en la DB: %s\n", hashedPassword)

	// --- SIMULACIÓN DEL LOGIN ---
	fmt.Println("\n--- SIMULACIÓN DEL LOGIN ---")

	// 1. El usuario introduce la contraseña para iniciar sesión.
	loginPassword := []byte("cambiar123")

	// 2. Se verifica la contraseña.
	// bcrypt.CompareHashAndPassword compara el hash almacenado con la contraseña ingresada.
	// bcrypt se encarga de extraer la sal del hash automáticamente y el costo.
	// La función utiliza la misma sal y el mismo costo que extrajo para hashear la contraseña en texto plano que el usuario ingresó.
	err = bcrypt.CompareHashAndPassword(hashedPassword, loginPassword)
	if err != nil {
		fmt.Println("Contraseña incorrecta. Acceso denegado.")
	} else {
		fmt.Println("¡Contraseña correcta! Acceso concedido.")
	}

	// --- EJEMPLO CON CONTRASEÑA INCORRECTA ---
	fmt.Println("\n--- EJEMPLO CON CONTRASEÑA INCORRECTA ---")
	wrongPassword := []byte("123*")

	err = bcrypt.CompareHashAndPassword(hashedPassword, wrongPassword)
	if err != nil {
		fmt.Println("Contraseña incorrecta. Acceso denegado.")
	} else {
		fmt.Println("¡Contraseña correcta! Acceso concedido.")
	}
}
