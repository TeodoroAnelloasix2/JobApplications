package main

import "fmt"

func main() {
	//declarar matriz--> var name_var [len]tipo_de_dato (var m [3]string)
	var matriz [5]int
	matriz[0] = 10
	fmt.Println(matriz)

	var matriz2 = [5]int{1, 2, 3, 4, 5} //declarar matriz, inicializar. Sabemos la longitud

	var matriz3 = [...]int{7, 8, 9} //con [...] indicamos longitud variable y dinamica

	fmt.Println(matriz2)

	//iterar matrices

	for i := range len(matriz3) {
		fmt.Println(matriz3[i])
	}
	for index, value := range matriz3 {
		fmt.Printf("indice: %d,valor: %d\n", index, value)
	}
	//matriz bidimensional var name [filas][columnas]tipodedato (var x [2][4]string{{fila1},{fila2},{fila3}})

	var m = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(m)
}
