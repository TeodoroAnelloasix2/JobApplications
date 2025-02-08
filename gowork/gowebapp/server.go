package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

var tmpl = template.Must(template.ParseFiles("./templates/index.html"))

func main() {

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", home)
	http.HandleFunc("/info", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "Host: ", req.Host)
		fmt.Fprintln(w, "URI: ", req.RequestURI)
		fmt.Fprintln(w, "Method: ", req.Method)
		fmt.Fprintln(w, "RemoteAddr: ", req.RemoteAddr)
	})

	http.HandleFunc("/products", products)
	http.HandleFunc("/redirect", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/products", 301)
	})
	http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Error...!", 501)
	})

	http.HandleFunc("/template", plantilla)
	http.HandleFunc("/cabeceras", cabecera)

	//http.ListenAndServe("0.0.0.0:8080", nil)
	http.ListenAndServeTLS(":8443", "./certs/cert.pem", "./certs/key.pem", nil)
}

func plantilla(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, struct{ Saludo string }{"Hola mundo!!!"})
}

func home(w http.ResponseWriter, r *http.Request) {

	html := "<html>"
	html += "<body>"
	html += "<h1>Hola Mundo!</h1>"
	html += "</html>"
	html += "</body>"
	w.Write([]byte(html))
}

func cabecera(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Test", "test1")

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Println(w, "{\"hola\":1}")

}

var productos []string

func products(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	add, okForm := r.Form["add"]
	if okForm && len(add) == 1 {
		productos = append(productos, string(add[0]))
		w.Write([]byte("Producto añadido correctamente"))
		return
	}
	prod, ok := r.URL.Query()["prod"]
	if ok && len(prod) == 1 {
		pos, err := strconv.Atoi(prod[0])

		if err != nil {
			return
		}
		html := "<html>"

		html += "<body>"
		html += "<h1>Producto: " + productos[pos] + "</h1>"
		html += "</html>"
		html += "</body>"

		w.Write([]byte(html))
		return
	}

	html := "<html>"
	html += "<body>"
	html += "<h1>Our Products:" + strconv.Itoa(len(productos)) + "</h1>"
	html += "</html>"
	html += "</body>"

	w.Write([]byte(html))
}
