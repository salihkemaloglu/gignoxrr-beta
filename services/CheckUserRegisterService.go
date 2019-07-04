package service

import (
	"context"
	"fmt"

	helper "github.com/salihkemaloglu/gignoxrr-beta-001/helpers"
	inter "github.com/salihkemaloglu/gignoxrr-beta-001/interfaces"
	gigxRR "github.com/salihkemaloglu/gignoxrr-beta-001/proto"
	repo "github.com/salihkemaloglu/gignoxrr-beta-001/repositories"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

//CheckUserToRegisterService ...
func CheckUserToRegisterService(ctx context.Context, req *gigxRR.CheckUserToRegisterRequest) (*gigxRR.CheckUserToRegisterResponse, error) {
	userLang := "en"
	if headers, ok := metadata.FromIncomingContext(ctx); ok {
		if headers["languagecode"] != nil {
			userLang = headers["languagecode"][0]
		}
	}
	lang := helper.DetectLanguage(userLang)

	userData := req.GetUser()
	user := repo.User{
		Username: userData.GetUsername(),
		Email:    userData.GetEmail(),
	}

	var userOp inter.IUserRepository = user
	if err := userOp.CheckUser(); err == nil {
		return nil, status.Errorf(
			codes.AlreadyExists,
			fmt.Sprintf(helper.Translate(lang, "already_created_account")+user.Username),
		)
	}

	return &gigxRR.CheckUserToRegisterResponse{
		GeneralResponse: &gigxRR.GeneralResponse{
			Message:            "Username or email ok ",
			IsEmailSuccess:     true,
			IsOperationSuccess: true,
		},
	}, nil
}
