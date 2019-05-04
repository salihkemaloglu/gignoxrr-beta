package controller

import (
	"time"
	"fmt"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/metadata"
	"github.com/salihkemaloglu/gignox-rr-beta-001/proto"
	val "github.com/salihkemaloglu/gignox-rr-beta-001/validations"
	repo "github.com/salihkemaloglu/gignox-rr-beta-001/repositories"
	inter "github.com/salihkemaloglu/gignox-rr-beta-001/interfaces"
	helper "github.com/salihkemaloglu/gignox-rr-beta-001/services"
)
const (
	timeFormat = "2006-01-02 15:04:05"
)

func  RegisterController(ctx_ context.Context, req_ *gigxRR.RegisterUserRequest) (*gigxRR.RegisterUserResponse, error) {
	userLang :="en"
	if headers, ok := metadata.FromIncomingContext(ctx_); ok {
		userLang = headers["language"][0]
	}
	lang := helper.DetectLanguage(userLang)

	userData := req_.GetUser();
	t := time.Now().UTC()
	user := repo.User {
		Name: userData.GetName(),
		Surname: userData.GetSurname(),
		Username: userData.GetUsername(),
		Email: userData.GetEmail(),
		Password: userData.GetPassword(),
		CreatedDate: t.Format(timeFormat),
		UpdatedDate: t.Format(timeFormat),
		TotalSpace: 100,
		LanguageCode: userLang,
	}
	
	var userOp inter.IUserRepository=user
	if err := userOp.CheckUser(); err ==nil  {
		return nil,status.Errorf(
			codes.AlreadyExists,
			fmt.Sprintf(helper.Translate(lang,"Already_Created_Account")+user.Username),
		)
	}
	if valResp := val.UserRegisterFieldValidation(user,lang); valResp != "ok" {
		return nil,status.Errorf(
			codes.FailedPrecondition,
			fmt.Sprintf(valResp),
		)
	}
	if dbResp := userOp.Insert(); dbResp != nil {
		return nil,status.Errorf(
			codes.Unimplemented,
			fmt.Sprintf(helper.Translate(lang,"Account_Insert_Error")+" :%v",dbResp.Error()),
		)
	}
	verificationCode,verErr:=helper.GenerateVerificationCodeService()
	if verErr !=nil {
		verificationCode = "134584"
	}
	isOk:=false
	emailResp:=helper.SendUserRegisterConfirmationMailService(user.Email,verificationCode,userLang);
	if emailResp != "ok" {
		isOk=true
	}

	return &gigxRR.RegisterUserResponse{
		GeneralResponse:&gigxRR.GeneralResponse{
			Message:emailResp,
			IsEmailSuccess:isOk,
			IsOperationSuccess:true,
		},
	}, nil
}