/*
Ejercicio: Calcule e imprima el área y el perímetro del triángulo
Crear un programa que solicite al usuario que ingrese los lados de un triángulo rectángulo y luego calcule e imprima el área y el perímetro del triángulo.
El programa debe:

	Solicitar al usuario que ingrese la longitud de los dos lados del triángulo rectángulo.
	Calcular la hipotenusa del triángulo usando el teorema de Pitágoras.
	Calcular el área del triángulo usando la fórmula base x altura / 2.
	Calcular el perímetro del triángulo sumando los lados.
	Imprimir el área y el perímetro del triángulo con dos decimales de precisión.
	El programa debe usar variables para almacenar los lados del triángulo, la hipotenusa, el área y el perímetro.

Además, se deben utilizar funciones del paquete Math de Go para calcular la raíz cuadrada y cualquier otro cálculo matemático necesario.

Ejemplo de entrada y salida:

	Ingrese lado 1: 3.5
	Ingrese lado 2: 4.2

	Área: 7.35
	Perímetro: 12.20
*/
package main

import "fmt"

var t *Triangulo_rectangulo

func main() {

	// Inicializar la estructura antes de usarla en Scan
	t = &Triangulo_rectangulo{}

	fmt.Println("Ingresa los valores de los lados del triangulo, ejemplo 3 4 ")
	fmt.Scan(&t.cateto, &t.cateto2)

	t.hipotenusa = t.Hpotnusa()

	fmt.Printf("La hipotenusa resultante es: %v\n", t.Hpotnusa())
	fmt.Printf("El area del triangulo es %v\n", t.Area())
	fmt.Printf("El perimetro del triangulo es %v \n", t.Perimetro())

}
