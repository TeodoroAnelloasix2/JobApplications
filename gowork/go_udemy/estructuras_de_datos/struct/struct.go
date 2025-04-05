package main

import "fmt"

type Persona struct {
	nombre string
	edad   int
	correo string
}

func main() {
	var p Persona

	p.nombre = "Alex"
	p.edad = 28
	p.correo = "alex@test.com"

	fmt.Println(p)

	p2 := Persona{"Anonymous", 28, "an@gmail.com"}
	fmt.Println(p2)

	p.edad = 30
	p2.edad = 35

	fmt.Printf("P %s tiene %d anni", p.nombre, p.edad)
	fmt.Println()
	//Punteros

	var x int = 10
	var puntero *int = &x

	fmt.Println(puntero)
	fmt.Println(&x)

	fmt.Println(x)

	editar(&x)
	fmt.Println(x)

	p2.saludar2()
	p.saludar()
}

func (p *Persona) saludar() {
	fmt.Println("Hola mi nombre es", p.nombre)
}

func (p2 *Persona) saludar2() {
	fmt.Println("Hola mi nombre es", p2.nombre)
}
func editar(x *int) {
	*x = 20

}
