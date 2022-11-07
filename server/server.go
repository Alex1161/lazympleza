package server

import "net/http"

func NewServer(port string) *http.Server {
	initRoutes()

	return &http.Server{
		Addr: port,
	}
}
