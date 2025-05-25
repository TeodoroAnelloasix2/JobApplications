package greetings

import (
	"errors"
	"fmt"

	"math/rand/v2"
)

func Hello(name string) (msg string, err error) {
	if name == "" {
		err = errors.New("nombre vacio")
		return "", err
	}
	msg = fmt.Sprintf(RandomFormat(), name)
	return msg, nil

}
func RandomFormat() (saludo string) {
	formats := []string{
		"Hola, %v. Bienvenido!",
		"Que bueno verte, %v",
		"Saludos, %v. Encantado de conocerte!",
	}
	saludo = formats[rand.IntN(len(formats))]
	return saludo
}

func Hellos(names []string) (map[string]string, error) {

	messages := make(map[string]string)
	for _, name := range names {
		msg, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = msg

	}
	return messages, nil
}
