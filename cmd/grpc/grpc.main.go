package main

import (
	"log"

	"github.com/michaelhendraw/stockbit_microservice/cmd/internal"
)

func main() {
	log.Println("Starting: Stockbit GRPC")

	// Read config
	config := internal.InitConfig()

	// Initialize everything here
	grpcServer := initializeService(config)

	log.Println("Running : Stockbit GRPC PORT", config.Server.GRPCPort)

	// Serve cron without HTTP TODO
	serve(grpcServer)
}

func initializeService(config internal.Config) *internal.GRPCServer {

	// Get all available usecases
	ucase := internal.GetUsecase(&config)

	// Initialize GRPC Server
	grpcServer := initGRPC(&config, ucase)

	return grpcServer
}
