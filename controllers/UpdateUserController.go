package controllers

import (
	"context"
	"fmt"

	gigxRR "github.com/salihkemaloglu/gignoxrr-beta-001/proto"
)

//UpdateUser ...
func (s *Server) UpdateUser(ctx context.Context, req *gigxRR.UpdateUserRequest) (*gigxRR.UpdateUserResponse, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("There's something wrong:", err)
		}
	}()
	return nil, nil
}
