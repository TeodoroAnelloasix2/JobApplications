package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

func conexionBD() (conexion *sql.DB) {
	Driver := "mysql"
	Usuario := "teo"
	Contrasenia := "system"
	NombreDB := "sistema"

	conexion, err := sql.Open(Driver, Usuario+":"+Contrasenia+"@tcp(127.0.0.1)/"+NombreDB)
	if err != nil {
		panic(err.Error())
	}
	return conexion
}

var plantillas = template.Must(template.ParseGlob("./templates/*"))

func main() {
	http.HandleFunc("/", inicio)
	http.HandleFunc("/crear", Crear)
	http.HandleFunc("/insert", insertar)
	http.HandleFunc("/borrar", Borrar)
	http.HandleFunc("/editar", Editar)
	http.HandleFunc("/actualizar", Actualizar)
	log.Println("Servidor corriendo...")
	http.ListenAndServe(":8080", nil)
}

func Borrar(w http.ResponseWriter, r *http.Request) {
	idEmpleado := r.URL.Query().Get("id")
	fmt.Println(idEmpleado)
	conexionEstablecida := conexionBD()

	borrarRegistro, err := conexionEstablecida.Prepare("Delete from Empleados where id=?")

	if err != nil {
		panic(err.Error())
	}

	borrarRegistro.Exec(idEmpleado)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func insertar(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		nombre := r.FormValue("nombre")
		correo := r.FormValue("correo")

		//.Prepare(query) necesita de .exec() para que se ejecute
		conexionEstablecida := conexionBD()

		insertarRegistros, err := conexionEstablecida.Prepare("INSERT INTO Empleados(nombre,correo) values(?,?)")
		if err != nil {
			panic(err.Error())
		}
		insertarRegistros.Exec(nombre, correo)
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}

type Empleado struct {
	Id     int
	Nombre string
	Correo string
}

func inicio(w http.ResponseWriter, r *http.Request) {
	//plantillas.ExecuteTemplate(w, "inicio", nil)

	conexionEstablecida := conexionBD()
	//.Query se ejecuta directamente var,err:=conexionEstablecida.query(query)   var contendrá el resultado

	selectRegistro, err := conexionEstablecida.Query("SELECT * FROM Empleados")

	if err != nil {
		panic(err.Error())
	}
	empleado := Empleado{}
	arregloEmpleado := []Empleado{}

	for selectRegistro.Next() {
		var id int
		var nombre string
		var correo string
		err = selectRegistro.Scan(&id, &nombre, &correo)
		if err != nil {
			panic(err.Error())
		}
		empleado.Id = id
		empleado.Nombre = nombre
		empleado.Correo = correo

		//nueva_lista=append(lista_existente,elemento_a_agregar)
		//array=append(array,elemento_new)

		arregloEmpleado = append(arregloEmpleado, empleado)

	}

	plantillas.ExecuteTemplate(w, "inicio", arregloEmpleado)
	//fmt.Println(arregloEmpleado)

}

func Crear(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "crear", nil)
}
func Editar(w http.ResponseWriter, r *http.Request) {

	idEmpleado := r.URL.Query().Get("id")
	conexionEstablecida := conexionBD()

	registro, err := conexionEstablecida.Query("Select * from Empleados where id=?", idEmpleado)
	if err != nil {
		panic(err.Error())
	}
	empleado := Empleado{}

	for registro.Next() {
		var id int
		var nombre string
		var correo string
		err = registro.Scan(&id, &nombre, &correo)
		if err != nil {
			panic(err.Error())
		}
		empleado.Id = id
		empleado.Nombre = nombre
		empleado.Correo = correo

		//nueva_lista=append(lista_existente,elemento_a_agregar)
		//array=append(array,elemento_new)

	}
	fmt.Println(empleado)
	plantillas.ExecuteTemplate(w, "editar", empleado)

}
func Actualizar(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		id := r.FormValue("id")
		nombre := r.FormValue("nombre")
		correo := r.FormValue("correo")

		conexionEstablecida := conexionBD()

		registroEditar, err := conexionEstablecida.Prepare("Update Empleados SET nombre=?,correo=? where id=?")
		if err != nil {
			panic(err.Error())
		}

		registroEditar.Exec(nombre, correo, id)
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}
