package service

import (
	
	"net"
	"fmt"
	"context"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/salihkemaloglu/gignox-rr-beta-001/proto"
)
func GetIpAddress(ctx context.Context) (*gigxRR.GetIpAddressResponse, error) {
	
	pr, ok := peer.FromContext(ctx)
	if !ok {
		return nil,status.Errorf(
			codes.Unimplemented,
			fmt.Sprintf("getClinetIP, invoke FromContext() failed"),
		)
	}
	if pr.Addr == net.Addr(nil) {
		return nil,status.Errorf(
			codes.Unimplemented,
			fmt.Sprintf("getClientIP, peer.Addr is nil"),
		)
	}	
	return &gigxRR.GetIpAddressResponse{IpAddress: pr.Addr.String()}, nil
}
