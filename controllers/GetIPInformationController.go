package controllers

import (
	"context"
	"fmt"

	gigxRR "github.com/salihkemaloglu/gignoxrr-beta-001/proto"
	serv "github.com/salihkemaloglu/gignoxrr-beta-001/services"
)

//GetIPInformation ...
func (s *Server) GetIPInformation(ctx context.Context, req *gigxRR.GetIPInformationRequest) (*gigxRR.GetIPInformationResponse, error) {

	fmt.Printf("RR service is working for GetIpAddess...Received rpc from client")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("There's something wrong:", err)
		}
	}()
	return serv.GetIPInformationService(ctx)
}
