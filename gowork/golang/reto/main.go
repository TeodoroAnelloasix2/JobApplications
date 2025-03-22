/*
Reto: Sistema de Notificaciones

Debes implementar un sistema de notificaciones donde haya diferentes tipos de notificaciones
(EmailNotification, SMSNotification) que implementen una interfaz común Notifier.
Requisitos:

Crea una interfaz Notifier con un método Send() que reciba un mensaje (string) y devuelva un string.
Implementa dos structs:

	EmailNotification que tenga un campo email (cadena).
	SMSNotification que tenga un campo phoneNumber (cadena).

Ambos structs deben implementar el método Send(), que retorne un mensaje indicando que se ha enviado una notificación por email o SMS.
Crea una función NotifyUser(n Notifier, message string) que reciba una instancia de Notifier y muestre el mensaje enviado.
En la función main(), crea instancias de EmailNotification y SMSNotification, y llama a NotifyUser() con ambas.

Ejemplo de salida esperada:

Enviando email a usuario@example.com: Hola, tienes un nuevo mensaje.
Enviando SMS a +34123456789: Hola, tienes un nuevo mensaje.
*/
package main

import (
	"fmt"
)

type Notifier interface {
	Send(string) string
}

type EmailNotification struct {
	email string
}

func (enot EmailNotification) Send(msg string) string {
	notificacion := fmt.Sprintf("Enviando email a %s: Hola, tienes un nuevo mensaje: %s", enot.email, msg)
	return notificacion
}

type SMSNotification struct {
	phoneNumber string
}

func (snot SMSNotification) Send(msg string) string {
	notificacion := fmt.Sprintf("Enviando SMS a +%s: Hola, tienes un nuevo mensaje: %s", snot.phoneNumber, msg)
	return notificacion
}
func NotifyUser(n Notifier, msg string) (notifica_enviada string) {
	notifica_enviada = fmt.Sprintf(n.Send(msg))
	return notifica_enviada
}
func main() {

	email1 := EmailNotification{
		email: "test1@test.com"}
	email2 := EmailNotification{
		email: "test2@test.com"}
	sms := SMSNotification{
		phoneNumber: "333333333"}

	fmt.Println(email1.Send("Test de email1"))
	fmt.Println(email2.Send("Test de email2"))
	fmt.Println(sms.Send("Test de sms"))
	fmt.Println(NotifyUser(email1, "test con Notify User"))
}
