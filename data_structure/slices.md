# Diferencias entre declarar slices con literal y con `make` en Go

En Go, tanto `[]string{"Barcelona", "Real Madrid", "Juventus", "PSG"}` como `make([]string, n)` crean **slices**, pero existen diferencias importantes en cómo se construyen y se inicializan.

---

## 🔹 Declaración literal de slice
```go
equipos := []string{"Barcelona", "Real Madrid", "Juventus", "PSG"}
```
- Creas un **slice literal** ya inicializado con esos elementos.  
- El **len** (longitud) es `4`.  
- La **cap** (capacidad) también es `4`.  
- Es inmutable en cuanto a tamaño inicial: no puedes declarar menos elementos de los que pusiste, aunque luego sí puedes hacer `append` para crecer dinámicamente.  

Ejemplo:
```go
fmt.Println(len(equipos)) // 4
fmt.Println(cap(equipos)) // 4
```

---

## 🔹 Declaración con `make`
```go
datos := make([]string, 4) // slice de 4 strings vacíos
```
- Creas un **slice vacío pero con espacio reservado**.  
- Todos los elementos están en su valor cero (`""` en el caso de string).  
- El **len** es `4`, porque ya tiene 4 elementos vacíos.  
- Puedes especificar capacidad distinta de longitud:
  ```go
  datos := make([]string, 0, 10) // longitud 0, capacidad 10
  ```
  Aquí no hay elementos aún (`len=0`), pero puedes hacer hasta 10 `append` sin que Go tenga que realocar memoria.  

---

## 🔹 Diferencia clave
- **Literal (`[]string{...}`)**:  
  → Lo usas cuando ya conoces los valores iniciales.  
- **`make([]T, len, cap)`**:  
  → Lo usas cuando quieres un slice vacío o con valores cero, pero con cierta longitud y/o capacidad predefinida para optimizar o controlar memoria.  

---

## 📌 Ejemplo práctico completo en Go

```go
package main

import "fmt"

func main() {
    // Declaración literal
    equipos := []string{"Barcelona", "Real Madrid", "Juventus", "PSG"}
    fmt.Println("equipos:", equipos)
    fmt.Println("len:", len(equipos), "cap:", cap(equipos))

    // Declaración con make
    datos := make([]string, 0, 5)
    fmt.Println("\ndatos:", datos)
    fmt.Println("len:", len(datos), "cap:", cap(datos))

    // Usando append
    datos = append(datos, "Milan")
    datos = append(datos, "Bayern")
    fmt.Println("\ndatos después de append:", datos)
    fmt.Println("len:", len(datos), "cap:", cap(datos))
}
```

---

✅ **Resumen:**  
- Ambos son slices (no arrays).  
- Con literal: slice ya lleno con valores.  
- Con `make`: slice inicializado con valores cero, con longitud/capacidad controlables.  
