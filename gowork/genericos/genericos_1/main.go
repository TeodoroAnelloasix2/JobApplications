package main

import "fmt"

func Sum[T int | float64](nums ...T) T {
	var total T

	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	fmt.Println(Sum(4, 5, 6, 7.8))
	//Opcionalmente se puede indicar que tipo de datos se van a pasar pero no es necesario Sum[int](1,2,3,4)

	var num1 integer = 100
	var num2 integer = 250
	fmt.Println(Sum2(num1, num2, 7, 8))

	fmt.Println(Sum3(num1, num2, 5, 99))
}

// Crear un tipo de dato
type integer int

// Se puede indicar el nuevo tipo a la funcion pero no es eficiente lo mejor es usar aproximacion de datos ~int
func Sum2[T int | integer](nums ...T) T {
	var tot T
	for _, num := range nums {
		tot += num
	}
	return tot
}

//Eficiente: ~int enteros y todos los que derivan de ello

func Sum3[T ~int](nums ...T) T {

	var tot T
	for _, num := range nums {
		tot += num

	}
	return tot
}
