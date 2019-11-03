package controllers

import (
	"context"
	"fmt"

	gigxRR "github.com/salihkemaloglu/gignoxrr-beta-001/proto"
	serv "github.com/salihkemaloglu/gignoxrr-beta-001/services"
)

//Register ...
func (s *Server) Register(ctx context.Context, req *gigxRR.RegisterUserRequest) (*gigxRR.RegisterUserResponse, error) {

	fmt.Printf("RR service is working for Register...Received rpc from client.\n")
	return serv.RegisterService(ctx, req)
}
