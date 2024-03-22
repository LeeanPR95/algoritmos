package main

import "fmt"

func main() {
	var p1, p2, p3, prom, examen, trabajoF, promFinal float64

	fmt.Println("Ingresa la primera nota: ")
	fmt.Scanln(&p1)

	fmt.Println("Ingresa la segunda nota: ")
	fmt.Scanln(&p2)

	fmt.Println("Ingresa la tercera nota: ")
	fmt.Scanln(&p3)

	prom = (float64(p1+p2+p3) / 3) * 0.55

	fmt.Printf("El promedio de las notas es %.2f \n ", prom)
	fmt.Println("")
	fmt.Println("Escribe la nota del examen: ")
	fmt.Scanln(&examen)

	fmt.Println("Escribe la nota del trabajo final: ")
	fmt.Scanln(&trabajoF)

	promFinal = float64(prom + (examen * 0.30) + (trabajoF * 0.15))
	fmt.Printf("El promedio final es: %.2f \n", promFinal)
}
