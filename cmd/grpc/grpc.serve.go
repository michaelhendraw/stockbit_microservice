package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/michaelhendraw/stockbit_microservice/cmd/internal"
)

func serve(grpcServer *internal.GRPCServer) {

	// subscribe to SIGINT signals
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)

	<-stopChan // wait for SIGINT

	// Shutdown grpc server
	if grpcServer != nil {
		grpcServer.Stop()
	}

	log.Println("Gracfully Stopped : Stockbit GRPC")

}
