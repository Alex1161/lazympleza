package server

import (
	"fmt"
	"net/http"
)

func initRoutes() {
	//http.HandleFunc("/", GetWorldCupsData)

	http.HandleFunc("/wc", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			fmt.Fprintf(w, "Respuesta")
		case "POST":
			//Post
		case "PATCH":
			//Patch
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Método no soportado")
			return

		}
	})
}
