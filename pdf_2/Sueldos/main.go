package main

import (
	"fmt"
)

func main() {
	const PAGOHR float32 = 1250.25
	var nombre string
	var horas int
	var sueldo float32

	fmt.Printf("El pago por horas es %.2f \n", PAGOHR)
	fmt.Println("Escribe tu nombre: ")
	fmt.Scanln(&nombre)

	fmt.Println("Escribe las horas trabajadas: ")
	fmt.Scanln(&horas)

	sueldo = PAGOHR * float32(horas)
	fmt.Printf("El sueldo del empleado %s es %.2f \n", nombre, sueldo)
}
