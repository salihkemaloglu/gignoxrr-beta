package controllers

import (
	"context"
	"fmt"

	gigxRR "github.com/salihkemaloglu/gignoxrr-beta-001/proto"
	serv "github.com/salihkemaloglu/gignoxrr-beta-001/services"
)

//ResetUserPassword ...
func (s *Server) ResetUserPassword(ctx context.Context, req *gigxRR.ResetUserPasswordRequest) (*gigxRR.ResetUserPasswordResponse, error) {

	fmt.Printf("RR service is working for ResetUserPassword...Received rpc from client.\n")
	return serv.ResetUserPasswordService(ctx, req)
}
