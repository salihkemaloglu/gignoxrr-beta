package controller

import (
	"time"
	"fmt"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/metadata"
	"github.com/salihkemaloglu/gignox-rr-beta-001/proto"
	db "github.com/salihkemaloglu/gignox-rr-beta-001/mongodb"
	val "github.com/salihkemaloglu/gignox-rr-beta-001/validation"
	repo "github.com/salihkemaloglu/gignox-rr-beta-001/repository"
	helper "github.com/salihkemaloglu/gignox-rr-beta-001/services"
)
const (
	timeFormat = "2006-01-02 15:04:05"
)

func  RegisterController(ctx context.Context, req *gigxRR.RegisterUserRequest) (*gigxRR.RegisterUserResponse, error) {
	userLang :="en"
	if headers, ok := metadata.FromIncomingContext(ctx); ok {
		userLang = headers["language"][0]
	}
	lang := helper.DetectLanguage(userLang)

	userData := req.GetUser();
	t := time.Now().UTC()
	user := db.User {
		Name: userData.GetName(),
		Surname: userData.GetSurname(),
		Username: userData.GetUsername(),
		Email: userData.GetEmail(),
		Password: userData.GetPassword(),
		CreatedDate: t.Format(timeFormat),
		UpdatedDate: t.Format(timeFormat),
		TotalSpace: 100,
		LanguageType: userLang,
	}
	
	if valResp := val.UserRegisterFieldValidation(user,lang); valResp != "ok" {
		return nil,status.Errorf(
			codes.FailedPrecondition,
			fmt.Sprintf(valResp),
		)
	}

	var userOp repo.UserRepository=user
	if err := userOp.CheckUser(); err ==nil  {
		return nil,status.Errorf(
			codes.AlreadyExists,
			fmt.Sprintf(helper.Translate(lang,"Already_Created_Account")+user.Username),
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