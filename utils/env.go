package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func readFromSystemEnv() {
	// Lee las variables directamente del entorno del sistema.
	testMongoURI := os.Getenv("TEST_MONGODB_URI")
	fmt.Println("------- Variables del Sistema -------")
	fmt.Println("TEST_MONGODB_URI:", testMongoURI)
	fmt.Println("-------------------------------------")
}

// readFromDotEnv carga las variables desde un archivo .env.
func readFromDotEnv() {
	// por defecto busca un archivo llamado .env en el directorio actual.
	err := godotenv.Load("utils/config.env")
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %s", err)
	}

	// Ahora puedes leer las variables como si estuvieran en el entorno del sistema.
	mongoURI := os.Getenv("MONGODB_URI")

	fmt.Println("------- Variables de .env -------")
	fmt.Println("MONGODB_URI:", mongoURI)
	fmt.Println("---------------------------------")
}

func Env() {
	//readFromSystemEnv()
	readFromDotEnv()
}
