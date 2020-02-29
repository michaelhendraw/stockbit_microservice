package internal

import (
	"github.com/michaelhendraw/stockbit_microservice/internal/usecase"
	"google.golang.org/grpc"
)

// Config struct
type Config struct {
	// Main
	Server ServerConfig

	// DB
	Redis   map[string]*RedisConfig
	Elastic map[string]*ElasticConfig

	// Vendor
	Agent AgentConfig
}

// Usecase struct
type Usecase struct {
	Search usecase.Search
}

// GRPCServer struct
type GRPCServer struct {
	Server  *grpc.Server
	Address string
}
