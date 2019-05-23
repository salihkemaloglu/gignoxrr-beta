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
		userLang = headers["languagecode"][0]
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
			fmt.Sprintf(helper.Translate(lang,"already_created_account")+user.Username),
		)
	}
	if valResp := val.UserRegisterFieldValidation(user,lang); valResp != "ok" {
		return nil,status.Errorf(
			codes.FailedPrecondition,
			fmt.Sprintf(valResp),
		)
	}
    user.Password = helper.EncryptePassword(user.Password)
	if dbResp := userOp.Insert(); dbResp != nil {
		return nil,status.Errorf(
			codes.Aborted,
			fmt.Sprintf(helper.Translate(lang,"account_insert_error")+" :%v",dbResp.Error()),
		)
	}
	token,tokenErr:=helper.CreateTokenEndpointService(user)
	isTokenSuccess:=true
	if tokenErr != nil{
    	isTokenSuccess=false
		token=tokenErr.Error()
	}
	
	_,err:=helper.SendUserRegisterConfirmationMailService(user.Email,"register",token,userLang);
	if err != nil {
		return &gigxRR.RegisterUserResponse{
			GeneralResponse:&gigxRR.GeneralResponse{
				Message:fmt.Sprintf(helper.Translate(lang,"email_send_error")+" :%v",err.Error()),
				Token:token,
				IsEmailSuccess:false,
				IsOperationSuccess:false,
				IsTokenSuccess:isTokenSuccess,
			},
		}, nil
	}
	return &gigxRR.RegisterUserResponse{
		GeneralResponse:&gigxRR.GeneralResponse{
			Message:"register",
			Token:token,
			IsEmailSuccess:true,
			IsOperationSuccess:true,
			IsTokenSuccess:isTokenSuccess,
		},
	}, nil
}