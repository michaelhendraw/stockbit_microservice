package internal

import (
	"log"

	"github.com/valyala/fasthttp/reuseport"
)

// Start GRPC server
func (s *GRPCServer) Start() error {
	l, err := reuseport.Listen("tcp4", s.Address)
	if err != nil {
		log.Println("func Start", err)
		return err
	}

	return s.Server.Serve(l)
}

// Stop GRPC server
func (s *GRPCServer) Stop() {
	s.Server.GracefulStop()
}
