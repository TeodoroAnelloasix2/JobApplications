package main

import (
	"fmt"
)

func main() {

	myslice := [3]string{"hola", "hola2", "hola3"}

	fmt.Println(myslice[0])

	var myslice2 []string

	myslice2 = append(myslice2, myslice[2])

	fmt.Println(myslice2[0])

}
