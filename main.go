package main

import (
	"context"
	"lazympleza/decisionTree"
	"lazympleza/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	predictions := decisionTree.CreateDecisionTree()

	ctx := context.Background()

	serverDoneChan := make(chan os.Signal, 1)
	signal.Notify(serverDoneChan, os.Interrupt, syscall.SIGTERM)

	svr := server.NewServer(":8080", predictions)

	go func() {
		err := svr.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	log.Println("Inicio el servidor")

	<-serverDoneChan

	svr.Shutdown(ctx)
	log.Println("Finalizo el servidor")
}
