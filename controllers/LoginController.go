package controllers

import (
	"context"
	"fmt"

	"github.com/patrickmn/go-cache"
	gigxRR "github.com/salihkemaloglu/gignoxrr-beta-001/proto"
	serv "github.com/salihkemaloglu/gignoxrr-beta-001/services"
)

var c *cache.Cache

//Login ...
func (s *Server) Login(ctx context.Context, req *gigxRR.LoginUserRequest) (*gigxRR.LoginUserResponse, error) {

	fmt.Printf("RR service is working for Login...Received rpc from client.\n")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("There's something wrong:", err)
		}
	}()
	return serv.LoginService(ctx, req, c)
}
