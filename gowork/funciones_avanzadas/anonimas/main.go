package main

import "fmt"

func main() {
	//Funcion anonima 1 metodo
	func() {
		fmt.Println("Hola desde anonima 1")
	}()

	//Funcion anonima 2 metodo
	anonima := func(nombre string) string {

		return fmt.Sprintf("Hola desde anonima 2 %s\n", nombre)

	}

	componerSaludos(anonima, "Teo")

	//Duplicar y triplicar son funciones
	r1, r2 := aplicarMult(duplicar, 5), aplicarMult(triplicar, 5)

	fmt.Println(r1, r2)
}

func aplicarMult(f func(int) int, n int) int {
	return f(n)
}

// pasar una funcion como argumento
func componerSaludos(f func(string) string, nombre string) {

	fmt.Println(f(nombre))

}

// Segundo ejemplo de funciones como parametros de otras funciones
func duplicar(n int) int {
	return n * 2
}

func triplicar(n int) int {
	return n * 3
}
