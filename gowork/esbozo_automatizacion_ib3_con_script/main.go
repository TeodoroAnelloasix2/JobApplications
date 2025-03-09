package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {
	//variables
	separador := "/"
	repo_seguros := []string{"/Test/", "/test2/", "/JobApplications/"} //Lista de repositorios seguros a comprobar

	// //parametros del programa --url http://ejemplo.com --path  /ruta/almacenaje
	var p_url = flag.String("url", "", "Url del archivo a descargar")
	var p_path = flag.String("path", "./", "Ruta donde almacenar el archivo")

	flag.Parse()

	path := *p_path
	url := *p_url

	if path == "./" {
		fmt.Println("Ruta por defecto seleccionada. Se puede definir una ruta con --path /tu/ruta/deseada ")
	}
	if url == "" {
		log_a_registrar := "Provee una url: --url http://ejemplo.com"
		dejar_constancia_en_log(path, []byte(log_a_registrar))
		log.Fatal("Error" + log_a_registrar)
	}
	path, err := crear_estructura_directorios(separador, path)
	if err != nil {
		log_a_registrar := fmt.Appendf(nil, "%v\n", err)
		dejar_constancia_en_log(path, log_a_registrar)
		log.Fatal(string(log_a_registrar))
	}

	for _, repo := range repo_seguros {
		ok, _ := regexp.MatchString(repo, url)
		if !ok {
			continue
		}

		nombre_archivo, _, err := archivo(path, separador, url)
		if err != nil {
			log_a_registrar := fmt.Appendf(nil, "%v\n", err)
			dejar_constancia_en_log(path, log_a_registrar)
			log.Fatalf("%s No se pudo crear archivo: %s", log_a_registrar, nombre_archivo)
		}
		break
	}

	dejar_constancia_en_log(path, []byte("Programa terminado con éxito!"))
}

func archivo(path, separador, url string) (nombre_archivo string, resp *http.Response, err error) {

	resp, err = http.Get(url)
	if err != nil {
		log_a_registrar := fmt.Appendf(nil, "%v\n", err)

		dejar_constancia_en_log(path, log_a_registrar)
		return "", nil, err
	}
	defer resp.Body.Close()
	nombre_archivo = sacar_nombre_archvio(separador, url)
	archivo_salida, err := os.Create(path + nombre_archivo)
	if err != nil {
		log_a_registrar := fmt.Appendf(nil, "%v\n", err)
		dejar_constancia_en_log(path, log_a_registrar)
		return "", nil, err
	}
	defer archivo_salida.Close()
	_, err = io.Copy(archivo_salida, resp.Body)
	return nombre_archivo, resp, err
}

func sacar_nombre_archvio(separador, url string) (nombre_archvio string) {

	separados := strings.Split(url, separador)
	ultimo_elemento := len(separados) - 1

	nombre_archvio = separados[ultimo_elemento]
	return nombre_archvio
}

func crear_estructura_directorios(separador, path string) (new_path string, err error) {
	script_ib3 := "/script_ib3"
	estructura := []string{script_ib3 + "/success", script_ib3 + "/error"}
	for _, dir := range estructura {
		err = os.MkdirAll(path+"/"+dir, 0766)

	}
	new_path = path + script_ib3 + separador
	return new_path, err
}

func dejar_constancia_en_log(path string, log []byte) {
	var file_error string
	a := strings.Split(path, "/")

	if a[len(a)-1] == "/" {
		file_error = path + "log_error_go.txt"
	} else {
		file_error = path + "/log_error_go.txt"
	}

	f, err := os.OpenFile(file_error, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0766)
	if err != nil {
		fmt.Printf("Error al abrir el log %s\n", err)
		return
	}
	defer f.Close()

	if _, err := f.Write(append(log, '\n')); err != nil {
		fmt.Printf("Error al escribir en el log: %s\n", err)
	}
}

//TO DO
//func transferir_archivo(host,script)(err){ TO DO }
