package client_test

import (
	"context"
	"testing"

	. "github.com/michaelhendraw/stockbit_microservice/grpc/client"
	pb "github.com/michaelhendraw/stockbit_microservice/grpc/proto"
)

func TestGetClientWithInterceptor(t *testing.T) {
	// t.Skip()
	o := &Options{
		Address:         ":9898",
		WithInterceptor: true,
	}
	c, err := GetClient(o)
	t.Logf("%+v %+v", c, err)

}

func TestClose(t *testing.T) {
	// t.Skip()
	o := &Options{
		Address: ":9898",
	}
	c, _ := GetClient(o)
	if c != nil {
		c.Close()
		t.Logf("%+v", c)
	}

}

func TestGetSomething(t *testing.T) {
	o := &Options{}
	o.Address = ":9898"
	c, err := GetClient(o)
	req := pb.SearchRequest{SearchWord: "batman", Pagination: 1}
	resp, err := c.Search(context.Background(), &req)
	t.Logf("%+v %+v", resp, err)

}
