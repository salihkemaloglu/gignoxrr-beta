package controllers

import (
	"context"
	"fmt"

	gigxRR "github.com/salihkemaloglu/gignoxrr-beta-001/proto"
	serv "github.com/salihkemaloglu/gignoxrr-beta-001/services"
)

//CheckUserToRegister ...
func (s *Server) CheckUserToRegister(ctx context.Context, req *gigxRR.CheckUserToRegisterRequest) (*gigxRR.CheckUserToRegisterResponse, error) {

	fmt.Printf("RR service is working for CheckUserToRegister...Received rpc from client.\n")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("There's something wrong:", err)
		}
	}()
	return serv.CheckUserToRegisterService(ctx, req)
}
