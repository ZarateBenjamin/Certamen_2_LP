package main

import "fmt"

// Función para sumar dos números enteros
func sumar(a int, b int) int {
	return a + b
}

// Función principal del programa
func main() {
	// Imprimir un mensaje de bienvenida
	fmt.Println("Bienvenido a Go!")

	// --- Variables y Tipos de Datos ---

	// Declarar e inicializar variables
	var numero int = 10
	var texto string = "Hola Mundo"
	var flotante float64 = 3.14
	var booleano bool = true

	// Imprimir variables
	fmt.Println("Número:", numero)
	fmt.Println("Texto:", texto)
	fmt.Println("Flotante:", flotante)
	fmt.Println("Booleano:", booleano)

	// --- Funciones ---

	// Llamar a una función y almacenar el resultado
	resultado := sumar(5, 3)
	fmt.Println("El resultado de la suma es:", resultado)

	// --- Condicionales ---

	// Usar una estructura if-else para evaluar condiciones
	if numero > 5 {
		fmt.Println("El número es mayor a 5")
	} else if numero == 5 {
		fmt.Println("El número es igual a 5")
	} else {
		fmt.Println("El número es menor a 5")
	}

	// --- Listas (Slices) ---

	// Crear y manipular un slice (lista dinámica)
	lista := []int{10, 20, 30, 40, 50}
	lista = append(lista, 60) // Añadir un elemento al slice

	// Iterar sobre un slice
	for i, valor := range lista {
		fmt.Println("Índice:", i, "Valor:", valor)
	}

	// --- Mapas (Diccionarios) ---

	// Crear y manipular un mapa (diccionario)
	diccionario := make(map[string]int)
	diccionario["uno"] = 1
	diccionario["dos"] = 2
	diccionario["tres"] = 3

	// Acceder y modificar elementos en un mapa
	fmt.Println("El número tres:", diccionario["tres"])
	delete(diccionario, "dos") // Eliminar un elemento del mapa
	fmt.Println("Mapa después de eliminar 'dos':", diccionario)
}
