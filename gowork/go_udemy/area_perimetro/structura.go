package main

import "math"

type Triangulo_rectangulo struct {
	cateto     float32
	cateto2    float32
	hipotenusa float32
}

func (t *Triangulo_rectangulo) Area() (ar float32) {

	ar = (t.cateto * t.cateto2) / 2

	return ar
}

func (t *Triangulo_rectangulo) Perimetro() (p float32) {

	return t.cateto + t.cateto2 + t.hipotenusa
}

func (t *Triangulo_rectangulo) Hpotnusa() (h float32) {

	qrt_cateto1 := math.Pow(float64(t.cateto), 2)
	qrt_cateto2 := math.Pow(float64(t.cateto2), 2)

	c := qrt_cateto1 + qrt_cateto2

	h = float32(math.Sqrt(c))
	t.hipotenusa = h
	return h
}
