package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	googleGRPC "google.golang.org/grpc"
)

func TestGRPC(t *testing.T) {
	s := &GRPCServer{
		Server: &googleGRPC.Server{},
	}
	assert.Error(t, s.Start())
	assert.Panics(t, func() { s.Stop() })
}
