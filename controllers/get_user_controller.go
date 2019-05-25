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
	val "github.com/salihkemaloglu/gignox-rr-beta-001/validations"
	helper "github.com/salihkemaloglu/gignox-rr-beta-001/services"
)

func GetUserController(ctx_ context.Context, req_ *gigxRR.GetUserRequest) (*gigxRR.GetUserResponse, error) {
	userLang :="en"
	if headers, ok := metadata.FromIncomingContext(ctx_); ok {
		if headers["languagecode"] != nil {
			userLang = headers["languagecode"][0]
		}
	}
	lang := helper.DetectLanguage(userLang)
	username := req_.GetUsername();
	user := repo.User {
		Username:username,
	}
	if valResp := val.GetUserFieldValidation(user,lang); valResp != "ok" {
		return nil,status.Errorf(
			codes.FailedPrecondition,
			fmt.Sprintf(valResp),
		)
	}

	var userOp inter.IUserRepository = &user
	userResp,userErr := userOp.GetUserByUsername()
	if userErr != nil {
		return nil,status.Errorf(
			codes.NotFound,
			fmt.Sprintf(helper.Translate(lang,"get_user_invalid_user")+userErr.Error()),
		)
	}
	return &gigxRR.GetUserResponse {
		User:&gigxRR.User {
			Email:userResp.Email,
			Description:userResp.Description,
			ImagePath:userResp.ImagePath,
		},
	}, nil

}