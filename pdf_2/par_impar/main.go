package main

import (
	"fmt"
)

func main() {
	/* Elabora un algoritmo para leer un número 
	   y determinar si éste es par o impar
	*/

	// variables

	var num int
	fmt.Printf("Ingrese el numero \n")
	fmt.Scanln(&num)

	if num % 2 == 0 {
		fmt.Printf("%d es un número par", num)
	} else {
		fmt.Printf("%d es un número impar", num)
	}
	
}
