# Guía para Instalar Go

## Instalar Go en Windows

1. **Descargar el instalador**  
   - Ve a [la página oficial de Go](https://go.dev/dl/).
   - Descarga el archivo `.msi` para Windows.

2. **Ejecutar el instalador**  
   - Abre el archivo `.msi` descargado.
   - Sigue los pasos del instalador y acepta los valores predeterminados.

3. **Configurar las variables de entorno** (normalmente se configuran automáticamente).  
   - Para verificar, abre la terminal (`cmd`) y ejecuta:  
     ```sh
     go version
     ```
   - Si ves algo como `go version goX.X.X`, la instalación fue exitosa.

---

## Instalar Go en macOS

1. **Usar el instalador oficial**  
   - Ve a [la página de descargas de Go](https://go.dev/dl/).
   - Descarga el archivo `.pkg` y ábrelo.
   - Sigue los pasos del instalador.

2. **Usar Homebrew (opcional, recomendado)**  
   - Abre la terminal y ejecuta:  
     ```sh
     brew install go
     ```

3. **Verificar la instalación**  
   - En la terminal, ejecuta:  
     ```sh
     go version
     ```
   - Si aparece `go version goX.X.X`, todo está bien.

---

## Instalar Go en Linux (Ubuntu/Debian)

1. **Descargar Go**  
   - Ve a [la página oficial de Go](https://go.dev/dl/) y copia el enlace del último archivo `.tar.gz`.

2. **Instalar Go**  
   - Abre la terminal y ejecuta:  
     ```sh
     wget https://go.dev/dl/goX.X.X.linux-amd64.tar.gz
     sudo rm -rf /usr/local/go
     sudo tar -C /usr/local -xzf goX.X.X.linux-amd64.tar.gz
     ```

3. **Configurar el PATH**  
   - Agrega Go al PATH editando el archivo `~/.profile` o `~/.bashrc`:  
     ```sh
     echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.profile
     source ~/.profile
     ```

4. **Verificar la instalación**  
   ```sh
   go version
   ```

---

## ¡Listo para programar en Go!

Ahora puedes escribir código en Go. Para verificar que todo funciona, crea un archivo `hello.go` con este contenido:

```go
package main

import "fmt"

func main() {
    fmt.Println("¡Hola, Go!")
}
```

Para ejecutarlo, usa:

```sh
go run hello.go
```

