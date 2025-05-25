package main

import "fmt"

func main() {
	//los slices son listas de python3, son dinamicas y se puede agregar y eliminar elementos

	var myslice []int

	fmt.Println(myslice)

	diasemana := []string{"Domingo", "Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado"}

	fmt.Println(diasemana)

	diaslaborales := diasemana[1:6]
	fmt.Println(diaslaborales)
	//agregar lista:=append(lista,elemento agregar)

	diaslaborales = append(diaslaborales, "Sabado", "Domingo", "Otro dia")
	fmt.Println(diaslaborales)

	//eliminar un elemento de una lista o slices
	diaslaborales = append(diaslaborales[:5], diaslaborales[7:]...)
	fmt.Println(diaslaborales)

	rebanada1 := []int{1, 2, 3, 4, 5}

	rebanada2 := make([]int, 5)

	//copy(dest,source)
	copy(rebanada2, rebanada1)

	fmt.Println(rebanada2)
}
