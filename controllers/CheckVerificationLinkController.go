package controllers

import (
	"context"
	"fmt"

	gigxRR "github.com/salihkemaloglu/gignoxrr-beta-001/proto"
	serv "github.com/salihkemaloglu/gignoxrr-beta-001/services"
)

//CheckVerificationLink ...
func (s *Server) CheckVerificationLink(ctx context.Context, req *gigxRR.CheckVerificationLinkRequest) (*gigxRR.CheckVerificationLinkResponse, error) {

	fmt.Printf("RR service is working for CheckVerificationLink...Received rpc from client.\n")
	return serv.CheckVerificationLinkService(ctx, req)
}
