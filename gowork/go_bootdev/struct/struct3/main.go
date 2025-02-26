/*
Tarea

En Textio, una estructura user representa a un titular de cuenta, y un sender es simplemente un user
con algunos datos adicionales específicos de envío.
Un sender es un usuario que tiene un campo rateLimit, el cual indica cuántos mensajes puede enviar.

Corrige el error embebiendo la estructura adecuada dentro de la otra.
*/
package main

type sender struct {
	rateLimit int
	user
}

type user struct {
	name   string
	number int
}
