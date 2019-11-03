package controllers

import (
	"context"
	"fmt"

	gigxRR "github.com/salihkemaloglu/gignoxrr-beta-001/proto"
	serv "github.com/salihkemaloglu/gignoxrr-beta-001/services"
)

//GetUser ...
func (s *Server) GetUser(ctx context.Context, req *gigxRR.GetUserRequest) (*gigxRR.GetUserResponse, error) {

	fmt.Printf("RR service is working for GetUser...Received rpc from client.\n")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("There's something wrong:", err)
		}
	}()
	return serv.GetUserService(ctx, req)
}
