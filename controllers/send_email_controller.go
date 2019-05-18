package controller

import (
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

func  SendEmailController(ctx_ context.Context, req_ *gigxRR.SendEmailRequest) (*gigxRR.SendEmailResponse, error) {
	userLang :="en"
	if headers, ok := metadata.FromIncomingContext(ctx_); ok {
		userLang = headers["language"][0]
	}
	lang := helper.DetectLanguage(userLang)

	mailData := req_.GetGeneralRequest();
	user := repo.User {
		Email: mailData.GetEmailAddress(),
	}

	if valResp := val.SendMailFieldValidation(mailData.GetEmailAddress(),mailData.GetEmailType(),lang); valResp != "ok" {
		return nil,status.Errorf(
			codes.FailedPrecondition,
			fmt.Sprintf(valResp),
		)
	}

	var userOp inter.IUserRepository=user
	if err := userOp.CheckUser(); err != nil {
		return nil,status.Errorf(
			codes.NotFound,
			fmt.Sprintf(helper.Translate(lang,"User_Account_Not_Exist_Account")+user.Email),
		)
	}


	verificationCode,verErr:=helper.GenerateRandomStringURLService(128)
	if verErr !=nil {
		return nil,status.Errorf(
			codes.Aborted,
			fmt.Sprintf(helper.Translate(lang,"Generate_Password_Verification_Token_Error")+verErr.Error()),
		)
	}


	if mailData.GetEmailType() == "register" {
		return helper.SendUserRegisterConfirmationMailService(mailData.GetEmailAddress(),mailData.GetEmailType(),verificationCode,userLang)
		
	} else if mailData.GetEmailType() == "forgot" {
	  return	helper.SendUserForgotPasswordVerificationMailService(mailData.GetEmailAddress(),mailData.GetEmailType(),verificationCode,userLang)
		
	} else {
		return nil,status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf(helper.Translate(lang,"Unknown_Email_type")),
		)
	} 


}