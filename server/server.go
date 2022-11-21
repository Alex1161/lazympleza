package server

import (
	"net/http"
)

func NewServer(port string, tree chan string) *http.Server {
	initRoutes(tree)

	return &http.Server{
		Addr: port,
	}
}
