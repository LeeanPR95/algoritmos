package main

import "fmt"

func main() {
	/*
	* Escribe un algoritmo que lea los coeficientes
	* a, b, c, d, e, f
	* y calcule los valores de X e Y
	*/

	var a, b, c, d, e, f, x, y float32

	fmt.Print("Escriba el valor de a: ")
	fmt.Scanln(&a)

	fmt.Print("Escriba el valor de b: ")
	fmt.Scanln(&b)

	fmt.Print("Escriba el valor de c: ")
	fmt.Scanln(&c)

	fmt.Print("Escriba el valor de d: ")
	fmt.Scanln(&d)

	fmt.Print("Escriba el valor de e: ")
	fmt.Scanln(&e)

	fmt.Print("Escriba el valor de f: ")
	fmt.Scanln(&f)

	x = float32(((c * e) - (b * f)) / ((a * e) - (b * d)))
	fmt.Printf("El valor de X es %.2f \n", x)

	y = float32(((a * f) - (c * d)) / ((a * e) - (b * d)))
	fmt.Printf("El valor de Y es %.2f \n", y)
}
