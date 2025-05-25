package main

import "fmt"

func main() {
	colors := map[string]string{
		"rojo":  "#FF0000",
		"verde": "#00FF00",
		"azul":  "#0000FF",
	}

	fmt.Println(colors)
	fmt.Println(colors["rojo"])

	//agregar
	colors["negro"] = "#00000"

	fmt.Println(colors["negro"])

	//comprobar si existe llave valor

	valor, ok := colors["verde"]

	fmt.Println(valor, ok)

	valor, ok = colors["amarillo"]

	fmt.Println(valor, ok)
	//eliminar un elemento de un mapa
	delete(colors, "azul")
	for clave, value := range colors {
		fmt.Printf("Clave: %s, Valor %s\n", clave, value)
	}

}
