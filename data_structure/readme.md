
# 📌 Estructuras de Datos en Go

## 🔹 Definiciones Básicas

✅ **Estructuras de datos**: Son formas organizadas de almacenar y manipular información, como **arrays, slices, listas, pilas, colas, mapas y estructuras (structs)**.  

✅ **Array (Vector) en Go**: Es una colección de elementos del mismo tipo con **tamaño fijo**. 

Se define con:  
```go
var arr [3]int = [3]int{1, 2, 3}
```
📌 **Característica clave:** Su tamaño no puede cambiar después de la declaración.  

✅ **Slice en Go**: Es una estructura similar a un array, pero **dinámica y flexible** en tamaño. 

Se define con:  
```go
slice := []int{1, 2, 3}
```
📌 **Característica clave:** Puede crecer o reducirse dinámicamente.  

✅ **Mapas en Go**: Son estructuras de datos clave-valor que permiten almacenar y acceder a datos de manera eficiente. Se definen con:  
```go
mapa := make(map[string]int)
mapa["edad"] = 25
fmt.Println(mapa["edad"])
```
📌 **Característica clave:** Permiten acceso rápido a los valores a través de claves únicas.

---

## 🔹 Diferencia principal entre Array, Slice y Mapas

| **Característica**  | **Array** | **Slice** | **Mapa** |
|---------------------|----------|----------|----------|
| **Tamaño**         | Fijo      | Dinámico | Dinámico |
| **Flexibilidad**   | No puede cambiar de tamaño | Puede crecer o reducirse | Claves y valores dinámicos |
| **Uso recomendado**| Cuando se conoce el tamaño fijo | Para listas dinámicas | Para asociaciones clave-valor |

✅ **Conclusión**: En Go, los **slices** y **mapas** son más usados que los **arrays** debido a su flexibilidad y eficiencia. 🚀

