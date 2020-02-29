package main

import (
	"log"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/michaelhendraw/stockbit_microservice/cmd/internal"
	googleGRPC "google.golang.org/grpc"
)

func initGRPC(config *internal.Config, ucase *internal.Usecase) *internal.GRPCServer {

	// Init server
	s := googleGRPC.NewServer(
		googleGRPC.StreamInterceptor(
			grpc_middleware.ChainStreamServer(
				grpc_opentracing.StreamServerInterceptor(),
			),
		), // The last middleware must be Recovery

		googleGRPC.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_opentracing.UnaryServerInterceptor(),
			),
		),
	)

	// Create GRPC service
	grpcServer := &internal.GRPCServer{
		Server:  s,
		Address: ":" + config.Server.GRPCPort,
	}

	// Failed to open server connection
	if grpcServer == nil {
		log.Fatalln("[FATAL] Can't initialize GRPC")
	}

	// Start GRPC
	go func() {
		if err := grpcServer.Start(); err != nil {
			log.Fatalln("[FATAL] Can't start GRPC", err)
		}
	}()

	return grpcServer
}
