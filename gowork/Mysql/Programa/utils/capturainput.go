package utils

import (
	"clase1/modelos"
	"fmt"
)

func CapturaInputMenuPrincipal() (Eleccion int) {

	fmt.Printf("0 Para salir del programa\n1 Para listar todos los empleados\n2 Para listar el empleado con id especificado\n3 Para agregar un nuevo cliente\n")
	fmt.Scanln(&Eleccion)
	return Eleccion
}

func EligeId() (identificador int) {

	fmt.Printf("Elige el ID del empleado a mostrar\n")
	fmt.Scanln(&identificador)
	return identificador
}

func CapturarDatosCliente() *modelos.Cliente {
	NuevoCliente := modelos.Cliente{}
	fmt.Printf("Inserta los datos del cliente:\n")
	fmt.Println("Nombre: ")
	fmt.Scanln(&NuevoCliente.Nombre)
	fmt.Println("Email: ")
	fmt.Scanln(&NuevoCliente.Correo)
	fmt.Println("Telefono: ")
	fmt.Scanln(&NuevoCliente.Telefono)
	fmt.Println("Fecha Alta formato YYYY-MM-DD")
	fmt.Scanln(&NuevoCliente.FechaAlta)
	return &NuevoCliente
}
