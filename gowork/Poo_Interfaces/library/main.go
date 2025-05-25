package main

import (
	"fmt"
	"lib/animal"
	"lib/book"
)

func main() {

	var myBook = book.Book{
		Title:  "Moby Dick",
		Author: "Herman Melville",
		Pages:  300,
	}
	mybook2 := book.NewBook("Divina comedia", "Dante Alighieri", 1000)

	libropriv := book.NewPrivateBook("Se questo é un uomo", "Primo levi", 260)

	fmt.Println("#################################################################")
	myBook.PrintInfo()
	fmt.Println("#################################################################")
	mybook2.PrintInfo()
	fmt.Println(*libropriv)

	libropriv.SetTitle("Se questo è un uomo (special edition)")
	fmt.Println("#################################################################")
	fmt.Println(libropriv.GetTitle())
	fmt.Println("#################################################################")
	texbook := book.NewTexBook("Matematica", "Mondadori", "Casa mondadori", "Prima elementare", 80)
	texbook.PrintInfo()

	//Usar metodos de interface
	book.Print(texbook)
	book.Print(&myBook)
	book.Print(mybook2)

	fmt.Println("################### INTERFACES ANIMALES ##################################")

	miperro := animal.Perro{Nombre: "Fuffi"}
	migato := animal.Gato{Nombre: "Italia"}

	animal.HacerSonido(&miperro)
	animal.HacerSonido(&migato)

	animales := []animal.Animal{
		&animal.Gato{Nombre: "Max 1"},
		&animal.Gato{Nombre: "Max 2"},
		&animal.Perro{Nombre: "Max 3"},
		&animal.Perro{Nombre: "Max 4"},
	}
	for _, animal := range animales {
		animal.Sonido()
	}
}
