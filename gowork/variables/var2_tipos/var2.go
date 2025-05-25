package main

import (
	"fmt"
	"math"

	"rsc.io/quote" //go get rsc.io/quote
)

func main() {

	var intero_negativo int8 = -127
	var intero_positivo uint = 10

	fmt.Printf("%d %d"+" "+quote.Hello(), intero_negativo, intero_positivo)

	fmt.Printf("\nPara saber valor maximo math.tipo")
	fmt.Printf("\nValor maximo de (math.uint8) uint8: %v", math.MaxUint8)

	fmt.Println()

	var valueBool bool = true

	if valueBool {
		fmt.Printf("Cambiamos el valor de valueBool %v a %v", valueBool, !valueBool)
		valueBool = false
	}

	var a byte = 'a'
	texto := "hola"

	fmt.Printf("\nEl valor de a en ASCII es \t%v", a)
	fmt.Printf("\n el primer caracter de hola var[pos] es: %v", texto[0])

	var r rune = '🙃' //rune es de tipo unicode
	fmt.Println("\n", r)
}
