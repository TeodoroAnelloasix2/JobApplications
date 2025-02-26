/*
Vamos a mejorar la lógica de autenticación de Textio.
Almacenamos los datos de autenticación de nuestros usuarios dentro de una estructura (struct) llamada authenticationInfo.
Necesitamos un método que tome esos datos y devuelva una cadena de autorización básica.
El formato de la cadena debe ser:
Authorization: Basic USUARIO:CONTRASEÑA
Crea un método en la estructura authenticationInfo llamado getBasicAuth que devuelva la cadena con el formato especificado.

	type rect struct {
	  width int
	  height int
	}

// area has a receiver of (r rect)
// rect is the struct
// r is the placeholder

	func (r rect) area() int {
	  return r.width * r.height
	}

	var r = rect{
	  width: 5,
	  height: 10,
	}

fmt.Println(r.area())
50
*/
package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

func (a authenticationInfo) getBasicAuth() string {
	cadena := fmt.Sprintf("Authorization: Basic %s:%s", a.username, a.password)
	return cadena
}
