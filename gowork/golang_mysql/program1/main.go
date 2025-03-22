/*
Conectar a la base de datos mysql
export DBUSER='tu usuario aqui'
export DBPASS='tu password aqui'
*/
package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

// Propiedades del empleado
type Empleado struct {
	id        int
	nombre    string
	email     string
	sueldo    int
	funciones string
}

var trabajadores_lista []Empleado

func main() {

	database_name := "empleados"
	//tabla:="Trabajadores"

	//Propiedades de la conexion
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: database_name,
	}

	db, err := establecer_conexion(&cfg)
	if err != nil {
		log.Println("Error de tipo: ", err)
	}
	//fmt.Println(db)

	var empl Empleado
	trabajadores_lista, err = listar_todos_trabajdores(empl, db)
	if err != nil {
		err = fmt.Errorf("error en funcion listar todos trabajadores: %s", err)
		log.Println(err)
	}
	fmt.Printf("Trabajadores: \n %v \n", trabajadores_lista)

	fmt.Println()

	log.Println("Cerrando programa")
	defer db.Close()
}

func listar_todos_trabajdores(empl Empleado, db *sql.DB) (lista_empleados []Empleado, err error) {
	var rows *sql.Rows

	if db == nil {
		err = errors.New("error: db es nil")

		return lista_empleados, err
	}
	rows, err = db.Query("Select * FROM  Trabajadores")
	if err != nil {
		err = fmt.Errorf("error ejecutando la query: %s", err)
		return lista_empleados, err
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&empl.id, &empl.nombre, &empl.email, &empl.sueldo, &empl.funciones); err != nil {
			err = fmt.Errorf("error durante definicion struct: %s", err)
			return lista_empleados, err
		}
		lista_empleados = append(lista_empleados, empl)

	}
	if err = rows.Err(); err != nil {
		err = fmt.Errorf("error recorriendo rows.Next(): %s", err)
		return lista_empleados, err
	}

	return lista_empleados, nil
}
func establecer_conexion(cfg *mysql.Config) (db *sql.DB, err error) {

	status := "up"

	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		err = fmt.Errorf("error durante conexion base de datos: %v", err)

	}

	if err := db.Ping(); err != nil {
		status = "down"
		err = fmt.Errorf("error durante ping a  base de datos: %v", err)
		return db, err
	}
	log.Println(status)
	fmt.Println()

	return db, err
}
