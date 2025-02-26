/*
 * Realiza una conexión desde el lenguaje que hayas seleccionado a la siguiente base de datos MySQL:
 * - Host: mysql-5707.dinaserver.com
 * - Port: 3306
 * - User: mouredev_read
 * - Password: mouredev_pass
 * - Database: moure_test
 *
 * Una vez realices la conexión, lanza la siguiente consulta e imprime el resultado:
 * - SELECT * FROM `challenges`
 *
 * Se pueden usar librerías para realizar la lógica de conexión a la base de datos.
 */

/*
mysql -u user -p -h host para conectarse:

mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| moure_test         |
+--------------------+
2 rows in set (0,09 sec)
mysql> use moure_test
# Reading table information for completion of table and column names
# You can turn off this feature to get a quicker startup with -A
Database changed
mysql> show tables;
+----------------------+
| Tables_in_moure_test |
+----------------------+
| challenges           |
+----------------------+
1 row in set (0,02 sec)
mysql> desc challenges;
+------------+--------------+------+-----+---------+----------------+
| Field      | Type         | Null | Key | Default | Extra          |
+------------+--------------+------+-----+---------+----------------+
| id         | int(11)      | NO   | PRI | NULL    | auto_increment |
| name       | varchar(100) | NO   |     | NULL    |                |
| difficulty | varchar(10)  | NO   |     | NULL    |                |
| date       | date         | NO   |     | NULL    |                |
+------------+--------------+------+-----+---------+----------------+
4 rows in set (0,02 sec)
mysql>
*/
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func connectBD() (conexion *sql.DB) {

	Driver := "mysql"
	Usuario := "mouredev_read"
	Constrasenia := "mouredev_pass"
	NombreBD := "moure_test"

	conexion, err := sql.Open(Driver, Usuario+":"+Constrasenia+"@tcp(mysql-5707.dinaserver.com:3306)/"+NombreBD)
	if err != nil {
		panic(err.Error())
	}
	//fmt.Println("conexión estalecida")
	return conexion
}

func main() {

	conn := connectBD()

	//fmt.Println(conn)
	ejec_query(conn)
}

func ejec_query(conexionEstablecida *sql.DB) {

	//conexionEstablecida := connectBD()

	defer conexionEstablecida.Close()

	query_select, err := conexionEstablecida.Query("SELECT * FROM challenges")

	if err != nil {
		panic(err.Error())

	}
	for query_select.Next() {
		var (
			id         int32
			name       string
			difficulty string
			date       string
		)
		if err := query_select.Scan(&id, &name, &difficulty, &date); err != nil {
			log.Fatal(err)
		}
		fmt.Print("===================================")
		fmt.Println("")
		result := fmt.Sprintf("Id: %v, Name: %v, Difficulty: %v, Date: %v", id, name, difficulty, date)
		fmt.Print(result)
		fmt.Println("")

	}
}
