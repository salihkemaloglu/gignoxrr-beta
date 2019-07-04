package service

import (
	"context"
	"fmt"

	helper "github.com/salihkemaloglu/gignoxrr-beta-001/helpers"
	inter "github.com/salihkemaloglu/gignoxrr-beta-001/interfaces"
	gigxRR "github.com/salihkemaloglu/gignoxrr-beta-001/proto"
	repo "github.com/salihkemaloglu/gignoxrr-beta-001/repositories"
	val "github.com/salihkemaloglu/gignoxrr-beta-001/validations"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

//GetUserService ...
func GetUserService(ctx context.Context, req *gigxRR.GetUserRequest) (*gigxRR.GetUserResponse, error) {
	userLang := "en"
	if headers, ok := metadata.FromIncomingContext(ctx); ok {
		if headers["languagecode"] != nil {
			userLang = headers["languagecode"][0]
		}
	}
	lang := helper.DetectLanguage(userLang)
	username := req.GetUsername()
	user := repo.User{
		Username: username,
	}
	if valResp := val.GetUserFieldValidation(user, lang); valResp != "ok" {
		return nil, status.Errorf(
			codes.FailedPrecondition,
			fmt.Sprintf(valResp),
		)
	}

	var userOp inter.IUserRepository = &user
	userResp, userErr := userOp.GetUserByUsername()
	if userErr != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf(helper.Translate(lang, "get_user_invalid_user")+userErr.Error()),
		)
	}
	return &gigxRR.GetUserResponse{
		User: &gigxRR.User{
			Email:       userResp.Email,
			Description: userResp.Description,
			ImagePath:   userResp.ImagePath,
		},
	}, nil

}
