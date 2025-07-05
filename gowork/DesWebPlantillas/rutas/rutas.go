package rutas

import (
	"desweb1/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func CargarRutas() *mux.Router {
	mux := mux.NewRouter()
	mux.HandleFunc("/", handlers.HandleHome())
	mux.HandleFunc("/nosotros", handlers.SobreNosotros())
	//url http://srv.com/ruta/param1/param2 formato path
	mux.HandleFunc("/parametros/{id:.*}/{slug:.*}", handlers.Parametros())

	//url http://srver.com/hola?nombre='test' formato query string
	mux.HandleFunc("/paramtrosquerystring", handlers.QueryString())

	mux.HandleFunc("/estructura", handlers.HandlEstrucutra())
	ArchivosEstaticosMux(mux)
	return mux
}

func ArchivosEstaticosMux(mux *mux.Router) { //Modificamos el mux directamente con su puntero

	s := http.StripPrefix("/recursos/", http.FileServer(http.Dir("./recursos/")))
	mux.PathPrefix("/recursos/").Handler(s)

}
