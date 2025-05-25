package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func SaveContatsToFile(contacts []Contact) error {

	file, err := os.Create("contacts.json")
	if err != nil {
		return err
	}

	defer file.Close()
	//Escribe json en los archivos
	encoder := json.NewEncoder(file)
	err = encoder.Encode(contacts)
	if err != nil {
		return err
	}
	return nil
}

func loadContactsFromFile(contacts *[]Contact) error {

	file, err := os.Open("contacts.json")
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts)
	if err != nil {
		return err
	}
	return nil
}
func main() {

	var contactos []Contact

	err := loadContactsFromFile(&contactos)
	if err != nil {
		fmt.Println("Error al cargar los contactos: ", err)
	}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("=== Gestor de Contactos ===\n",
			"1. Agregar un contacto\n",
			"2. Mostrar los  contactos\n",
			"3. Salir\n",
			"Elige una opcion: \n")

		var option int

		if _, err = fmt.Scanln(&option); err != nil {

			log.Fatalln("Error al leer la opcion: ", err)

		}

		switch option {
		case 1:
			var c Contact

			fmt.Print("Nombre: ")
			c.Name, _ = reader.ReadString('\n')
			fmt.Print("Email: ")
			c.Email, _ = reader.ReadString('\n')
			fmt.Print("Telefon: ")
			c.Phone, _ = reader.ReadString('\n')
			contactos = append(contactos, c)
			if err := SaveContatsToFile(contactos); err != nil {
				fmt.Println("Error al guardar al contacto: ", err)
			}
		case 2:
			fmt.Println("===========================================")
			for index, contact := range contactos {
				fmt.Printf("%d. Nombre:%s Emails: %s Telefono: %s\n",
					index+1, contact.Name, contact.Email, contact.Phone)

			}
			fmt.Println("===========================================")
		case 3:
			return
		default:
			fmt.Println("Opcion invalida")
		}

	}

}
