package server

import (
	"lazympleza/memoize"
	"net/http"
)

func NewServer(port string, predictions memoize.MemoizedFunction) *http.Server {
	initRoutes(predictions)

	return &http.Server{
		Addr: port,
	}
}
