package main

import "fmt"

func Divide(op1, op2 int) {
	//Recover impide que panic interrumpa  el programa
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r) //Muestra mensaje de panico
		}
	}()
	Validar_dividedbyzero(op2)
	fmt.Println(op1 / op2)
}

func Validar_dividedbyzero(n int) {
	if n == 0 {
		panic("No puedes dividir por cero") //Interrumpe un programa de forma forzada
	}
}

func main() {
	Divide(100, 10)
	Divide(100, 0)

	Divide(100, 25)
}
