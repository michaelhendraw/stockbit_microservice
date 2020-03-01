package client

import (
	pb "github.com/michaelhendraw/stockbit_microservice/grpc/proto"
	"google.golang.org/grpc"
)

// Client for GRPC
type Client struct {
	conn *grpc.ClientConn
	grpc pb.NewStockbitClient
}

// Options for GRPC connection
type Options struct {
	Address         string
	WithInterceptor bool // set true to generate log
}
