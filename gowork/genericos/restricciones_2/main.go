package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Product[T uint | string] struct {
	Id    T
	Desc  string
	Price float32
}

// comparable es una palabra clave en golang que permite realizar comparaciones con tipos genericos T
func EsIncluido[T comparable](list []T, value T) bool {

	//return slices.Contains(list, value)
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

func Filter[T constraints.Ordered](list []T, callback func(T) bool) []T {
	result := make([]T, 0, len(list))

	for _, val := range list {
		if callback(val) {
			result = append(result, val)
		}
	}
	return result
}
func main() {
	var producto1 = Product[uint]{1, "zapatos", 50}
	var producto2 = Product[string]{"id-2", "bolso", 46}

	fmt.Println(producto1)
	fmt.Println(producto2)
	letras := []string{"a", "B", "c", "r"}
	numeros := []int{1, 2, 3, 4, 5}

	fmt.Println(EsIncluido(letras, "c"))
	fmt.Println(EsIncluido(letras, "f"))
	fmt.Println(EsIncluido(numeros, 1))
	fmt.Println(EsIncluido(numeros, 19))

	fmt.Println(Filter(numeros, func(value int) bool { return value > 3 }))
	fmt.Println(Filter(letras, func(value string) bool { return value > "b" }))
}
