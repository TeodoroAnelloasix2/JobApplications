package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ingrese su nombre y apellido: ")
	fmt.Scanln(&nombre, &apellido)
	fmt.Println("Ingrese su edad: ")
	fmt.Scanln(&edad)

	fmt.Printf("Me llamo %s  %s y tengo %d\n", nombre, apellido, edad)
	fmt.Printf("El tipo de la variable nombre es: %T\n", nombre)

}
