package controllers

import (
	"context"
	"fmt"

	gigxRR "github.com/salihkemaloglu/gignoxrr-beta-001/proto"
	serv "github.com/salihkemaloglu/gignoxrr-beta-001/services"
)

//SendEmail ...
func (s *Server) SendEmail(ctx context.Context, req *gigxRR.SendEmailRequest) (*gigxRR.SendEmailResponse, error) {

	fmt.Printf("RR service is working for SendMail...Received rpc from client.\n")
	return serv.SendEmailService(ctx, req)
}
