package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func main() {

	fmt.Println(Sumar(15, 3.5))
	fmt.Println(Sumar2[uint32](3, 4, 5))

}

// Crear restricciones con interface
// tipo number que tiene como contrato respetar los tipos contemplados
type Numbers interface {
	~float64 | ~uint | ~int
}

// La funcion acepta cualquier tipo de dato especificado por ---> [T type]
func Sumar[T Numbers](nums ...T) T {
	var tot T
	for _, num := range nums {
		tot += num
	}
	return tot
}

// https://pkg.go.dev/golang.org/x/exp/constraints Golang ofrece restricciones ya predefinidas en el paquete constraints.
func Sumar2[T constraints.Unsigned](nums ...T) T {
	var tot T
	for _, num := range nums {
		tot += num
	}
	return tot
}
