package main

import (
	"fmt"
	"log"

	"github.com/modulo/greetings"
)

func main() {
	testeo := "Teo"
	nombres := []string{"Alex", "Josef", "Mikol"}

	log.SetPrefix("greetings: ")
	log.SetFlags(0) //Estable ce bandera de formato que incluye info adicional al error (hora , fecha...)

	msg, err := greetings.Hello(testeo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)

	messages, err := greetings.Hellos(nombres)
	if err != nil {
		log.Fatal(err)
	}

	for _, n := range nombres {
		fmt.Println(messages[n])
	}
}
