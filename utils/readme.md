
## Gestión de Dependencias en Go

Go utiliza un sistema de **módulos** para gestionar las dependencias de los proyectos, asegurando que tu código sea reproducible y que las versiones de las librerías sean consistentes.

* **`go.mod`**: Este archivo define el módulo de tu proyecto y lista todas las dependencias directas e indirectas que tu proyecto necesita, junto con sus versiones específicas. Cuando ejecutas `go get` o importas un paquete en tu código, Go añade automáticamente la dependencia a este archivo. También incluye la **versión de Go** con la que se espera que se compile el proyecto (ej., `go 1.24`).

* **`go.sum`**: Complementa a `go.mod`. Contiene **sumas de verificación criptográficas** (hashes) de los módulos listados en `go.mod`. Esto proporciona una capa adicional de seguridad, ya que Go puede verificar la integridad de los módulos descargados, asegurándose de que no hayan sido alterados maliciosamente. Si los hashes no coinciden, Go te alertará.

* **`go mod tidy`**: Este comando es fundamental para mantener tus dependencias ordenadas. Realiza dos funciones principales:
    1.  **Añade dependencias faltantes**: Si has importado un nuevo paquete en tu código pero aún no está en `go.mod`, `go mod tidy` lo añade.
    2.  **Elimina dependencias no utilizadas**: Si eliminaste código que dependía de un paquete, `go mod tidy` lo eliminará de tu `go.mod` y `go.sum` si ya no es necesario.
    Es una buena práctica ejecutar `go mod tidy` regularmente para mantener tu archivo `go.mod` limpio y preciso.

### Para nuestro ejemplo instalamos:
 ```sh
 go get github.com/JimcostDev/jimcostdev-utils
 ```