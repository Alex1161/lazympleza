package server

import (
	"lazympleza/lazy"
	"net/http"
)

func NewServer(port string, predictions lazy.LazyFunction) *http.Server {
	initRoutes(predictions)

	return &http.Server{
		Addr: port,
	}
}
