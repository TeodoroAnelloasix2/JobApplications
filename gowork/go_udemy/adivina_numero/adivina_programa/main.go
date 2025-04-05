package main

import (
	"fmt"
	"math/rand"
)

var (
	aleatorio    int = rand.Intn(50)
	max_intentos int = 10
	intentados   int = 0
	num_jugador  int
)

func main() {
	Jugar()
}

func Jugar() {

	for intentados < max_intentos {
		fmt.Printf("Ingresa un numero del 1 al 50 (intentos restantes: %d)", max_intentos-intentados)
		intentados++
		fmt.Scanln(&num_jugador)

		if num_jugador == aleatorio {
			fmt.Println("Has ganado!")
			fmt.Printf("El numero es : %d\n", aleatorio)
			return
		} else if aleatorio < num_jugador {
			fmt.Println("El numero es menos al que has ingresado")
		} else if aleatorio > num_jugador {
			fmt.Println("El numero es mayor al que has ingresado")
		}
	}
	fmt.Println("Numero de intentos acabados! Has perdido pringado")
	fmt.Printf("El numero era el : %d\n", aleatorio)

}
