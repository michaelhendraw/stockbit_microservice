package grpc

import (
	"github.com/michaelhendraw/stockbit_microservice/internal/usecase"
)

// NewService ...
func NewService(
	searchUC usecase.Search,
) *GRPCService {
	return &GRPCService{
		NewSearchGRPC(searchUC),
	}
}
