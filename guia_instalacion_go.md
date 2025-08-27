# Guía de Instalación de Go

## Windows
1. Descarga el instalador `.msi` desde [go.dev/dl](https://go.dev/dl/)
2. Ejecuta el instalador y sigue los pasos
3. Verifica con `go version` en cmd

## macOS
**Opción 1 - Homebrew (recomendado):**
```sh
brew install go
```

**Opción 2 - Instalador oficial:**
1. Descarga el archivo `.pkg` desde [go.dev/dl](https://go.dev/dl/)
2. Ejecuta el instalador

## Linux
**Opción 1 - APT (Ubuntu/Debian):**
```sh
sudo apt update
sudo apt install golang-go
```

**Opción 2 - Instalación manual:**
```sh
# Descargar y extraer
wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz

# Agregar al PATH
echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc
source ~/.bashrc
```

## Verificación
Ejecuta en cualquier sistema:
```sh
go version
```

## Primer programa
Crea `hello.go`:
```go
package main

import "fmt"

func main() {
    fmt.Println("¡Hola, Go!")
}
```

Ejecuta con:
```sh
go run hello.go
```