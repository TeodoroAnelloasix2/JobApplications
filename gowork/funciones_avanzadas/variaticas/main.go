package main

import "fmt"

// Recibir n argumentos del mismo tipo
func sumar(numeros ...int) int {

	var tot int
	fmt.Printf("La variable que recibe n numeros (num ...int) es de tipo: %T y coniene - %v\n", numeros, numeros)
	for _, val := range numeros {
		tot += val
	}

	return tot
}

//Recibir n argumentos de diferentes tipo

func imprimirDatos(datos ...interface{}) {

	for _, val := range datos {
		fmt.Printf("Tipo: %T, Valor: %v\n", val, val)
	}
}

func main() {

	res := sumar(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(res)

	imprimirDatos(1, 2, "test", "gg", true, 1.7)
}
