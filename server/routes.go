package server

import (
	"fmt"
	"net/http"
)

func notSupportedMethod(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintf(w, "Método no soportado")
}

func initRoutes() {
	http.HandleFunc("/winner_between", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			notSupportedMethod(w, r)
			return
		}
		// Ejemplo de como catchear los parametros
		//log.Println(r.URL.Query()["var"])

		home := r.URL.Query()["home"][0]
		away := r.URL.Query()["away"][0]
		GetWinnerBetweet(w, r, home, away)
	})

	http.HandleFunc("/wc/cups", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			notSupportedMethod(w, r)
			return
		}
		GetWorldCupsData(w, r)
	})

	http.HandleFunc("/wc/matches", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			notSupportedMethod(w, r)
			return
		}
		GetWorldCupsMatches(w, r)
	})

	http.HandleFunc("/wc/players", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			notSupportedMethod(w, r)
			return
		}
		GetWorldCupsPlayers(w, r)
	})
}
