package main

import (
	"log"
	"os"
)

func main() {

	file, err := os.Create("hola.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_, err = file.Write([]byte("Hola, Soy anonymous"))

	if err != nil {
		log.Fatal(err)
	}

}
