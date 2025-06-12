## 1. El tipo `error` y `err != nil`

En Go el tipo de error es una interfaz incorporada:

```go
type error interface {
    Error() string
}
```

Cuando una función puede fallar, típicamente devuelve dos valores: el resultado y un error.

Para saber si ocurrió un error, simplemente compruebas:

```go
resultado, err := algunaFunción()
if err != nil {
    // ¡Hubo un error! Lo manejas aquí.
    fmt.Println("Error:", err)
    return
}
// Aquí err es nil y puedes usar resultado con seguridad.
```

**¿Por qué `err != nil`?**  
Porque en Go la ausencia de error se representa con `nil`; cualquier otro valor indica un fallo y suele traer un mensaje descriptivo.

---

## 2. `defer` para tareas de limpieza

La palabra clave `defer` retrasa la ejecución de una llamada hasta que la función actual termine (ya sea por retorno normal o por pánico). Muy útil para cerrar archivos, conexiones, desbloquear mutex, etc.

```go
func copiarArchivo(origen, destino string) error {
    f1, err := os.Open(origen)
    if err != nil {
        return err
    }
    // Garantiza que f1.Close() se ejecute al salir de la función
    defer f1.Close()

    f2, err := os.Create(destino)
    if err != nil {
        return err
    }
    defer f2.Close()

    _, err = io.Copy(f2, f1)
    return err
}
```

Aquí, aunque retor­nes antes por error, los `defer` se ejecutan y cierran los archivos correctamente.

---

## 3. `panic` y `recover` para emergencias

- `panic` detiene la ejecución normal y comienza a deshacer (“unwind”) las llamadas, ejecutando sus `defer`.
- `recover` solo funciona dentro de una función diferida (`defer`) y “recupera” el valor del pánico, evitando que el programa termine abruptamente.

```go
func protegido() {
    if r := recover(); r != nil {
        fmt.Println("¡Recuperado del pánico:", r)
    }
}

func ejemploPanic() {
    defer protegido()
    fmt.Println("Antes del pánico")
    panic("algo muy malo pasó")
    fmt.Println("Esto no se imprime")
}
```

Al llamar a `ejemploPanic()`, verás:

```text
Antes del pánico
¡Recuperado del pánico: algo muy malo pasó
```

**¿Cuándo usarlo?**  
Solo para situaciones realmente excepcionales de las que NO puedas recuperarte limpiamente con un error normal.

---

## 4. ¿Por qué no hay `try/catch` en Go?

1. **Claridad y explícito**  
   Al obligarte a comprobar cada `err != nil`, el código deja claro dónde pueden ocurrir fallos; no hay errores “ocultos”.

2. **Simplicidad en el lenguaje**  
   Go evita complejidad sintáctica y modelos de excepción: todo error es un valor más.

3. **Rendimiento predecible**  
   Las costosas operaciones de envío de stack traces solo ocurren en `panic`, no en el flujo normal.

---

## Resumen rápido

- **Errores son valores** (`error`) que siempre debes comprobar con `if err != nil`.  
- **`defer`** sirve para limpieza garantizada.  
- **`panic/recover`** son para emergencias y se combinan con `defer`.  
- **No hay `try/catch`** para mantener el lenguaje simple y el flujo de errores explícito.

¡Así de sencillo y poderoso es el manejo de errores en Go! 🚀
