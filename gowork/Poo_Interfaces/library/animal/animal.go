package animal

import "fmt"

type Animal interface {
	Sonido()
}

type Perro struct {
	Nombre string
}

func (p *Perro) Sonido() {
	fmt.Println(p.Nombre + " hace bau bau")
}

type Gato struct {
	Nombre string
}

func (g *Gato) Sonido() {
	fmt.Println(g.Nombre + " hace miao maio")
}

// Funcion que realiza el llamado segun el tipo de dato que recibe mientras cumpla con el tipo interface que espera recibir
func HacerSonido(an Animal) {
	an.Sonido()
}
