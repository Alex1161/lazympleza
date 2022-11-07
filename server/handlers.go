package server

import (
	"fmt"
	"net/http"
)

func GetWorldCupsData(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Aca tenes tus paises")
}
