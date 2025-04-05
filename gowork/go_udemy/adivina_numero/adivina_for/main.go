package main

import "fmt"

func main() {

	//Bucle infinito for{ code  } bucle con condicion for condition{ code  }
	var i int
	for i <= 10 {
		fmt.Println(i)
		i++
	} //for i:=1;i<=10;i++{ }

	for a := 1; a <= 10; a++ {
		if a == 4 {
			continue
		}
		fmt.Println(a)
	}
}
