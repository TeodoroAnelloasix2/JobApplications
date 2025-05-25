package main

import (
	"fmt"
)

const Pi float32 = 3.14 //Declaramos una constante a nivel de paquete
const (
	x = 10     //base 10
	y = 0b1010 //0b binario
	z = 0o12   //0o octal
	w = 0xFF   //Hexadecimal

)
const (
	Domingo = iota + 1 //Asigna valor sequencuial, empieza desde 0
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func main() {

	var (
		firstname, lastname string
		age                 int
	)
	firstname = "Alex"
	lastname = "kaley"
	age = 27

	fmt.Println(firstname, lastname, age)

	var coche, model, fecha = "seat", "ibiza", "2011"

	fmt.Printf("tengo un %s %s del %s", coche, model, fecha)

	fmt.Printf("\nEl valor de Pi es: %v", Pi)
	fmt.Println()
	// Imprimir valores numéricos en distintos formatos
	fmt.Println("Valores numéricos:")
	fmt.Printf("Decimal    : %d\n", x)
	fmt.Printf("Binario    : %b (El número %d en decimal es %b en binario)\n", y, x, x)
	fmt.Printf("Octal      : %o (El número %d en decimal es %o en octal)\n", z, x, x)
	fmt.Printf("Hexadecimal: %#X (El número %d en decimal es %#X en hexadecimal)\n", w, x, x)

	// Imprimir días de la semana
	fmt.Println("\nDías de la semana:")
	fmt.Printf("Domingo    : %d\n", Domingo)
	fmt.Printf("Lunes      : %d\n", Lunes)
	fmt.Printf("Martes     : %d\n", Martes)
	fmt.Printf("Miércoles  : %d\n", Miercoles)
	fmt.Printf("Jueves     : %d\n", Jueves)
	fmt.Printf("Viernes    : %d\n", Viernes)
	fmt.Printf("Sábado     : %d\n", Sabado)

}
