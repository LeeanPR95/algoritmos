package main

import "fmt"

func main() {
	// Algoritmo Evaluación
	// Descripcion: elaborar un algoritmo que obtenga e imprima el valor de Y a partir de
	// la ecuación
	// Y = 3*X^2 + 7X -15
	var x, y float32

	fmt.Print("Escribe el valor de X: ")
	fmt.Scanln(&x)

	y = float32((3 * x * x) + (7 * x) - 15)

	fmt.Printf("El valor de Y en la ecuación Y = 3X^2 + 7X - 15 es %.2f \n", y)
}
