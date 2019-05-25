package controller

import (
	"fmt"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/metadata"
	"github.com/salihkemaloglu/gignox-rr-beta-001/proto"
	repo "github.com/salihkemaloglu/gignox-rr-beta-001/repositories"
	inter "github.com/salihkemaloglu/gignox-rr-beta-001/interfaces"
	helper "github.com/salihkemaloglu/gignox-rr-beta-001/services"
)

func CheckUserToRegisterController(ctx_ context.Context, req_ *gigxRR.CheckUserToRegisterRequest) (*gigxRR.CheckUserToRegisterResponse, error) {
	userLang :="en"
	if headers, ok := metadata.FromIncomingContext(ctx_); ok {
		if headers["languagecode"] != nil {
			userLang = headers["languagecode"][0]
		}
	}
	lang := helper.DetectLanguage(userLang)

	userData := req_.GetUser();
	user := repo.User {
		Username: userData.GetUsername(),
		Email: userData.GetEmail(),
	}
	
	var userOp inter.IUserRepository=user
	if err := userOp.CheckUser(); err ==nil  {
		return nil,status.Errorf(
			codes.AlreadyExists,
			fmt.Sprintf(helper.Translate(lang,"already_created_account")+user.Username),
		)
	}

	return &gigxRR.CheckUserToRegisterResponse{
		GeneralResponse:&gigxRR.GeneralResponse{
			Message:"Username or email ok ",
			IsEmailSuccess:true,
			IsOperationSuccess:true,
		},
	}, nil
}