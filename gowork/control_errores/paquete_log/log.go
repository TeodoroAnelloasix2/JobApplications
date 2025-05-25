package main

import (
	"log"
	"os"
)

func main() {

	//Abrimos un archivo
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {

		log.Fatal(err)
	}
	defer file.Close()

	//Seteamos el output. Los registros se escribiràn en el archivo
	log.SetOutput(file)
	log.SetPrefix("Esto es un prefijo: ")
	log.Print("Soy un log")

}
