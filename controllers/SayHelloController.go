package controllers

import (
	"context"
	"fmt"

	gigxRR "github.com/salihkemaloglu/gignoxrr-beta-001/proto"
)

//Server ...
type Server struct {
}

//SayHello ...
func (s *Server) SayHello(ctx context.Context, req *gigxRR.HelloRequest) (*gigxRR.HelloResponse, error) {

	fmt.Printf("RR service is working for SayHello...Received rpc from client, message=%s\n", req.GetMessage())
	return &gigxRR.HelloResponse{Message: "Hello RR service is working..."}, nil
}
