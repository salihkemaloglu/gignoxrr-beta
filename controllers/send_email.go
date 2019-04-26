package controller

import (
	"fmt"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/metadata"
	"github.com/salihkemaloglu/gignox-rr-beta-001/proto"
	val "github.com/salihkemaloglu/gignox-rr-beta-001/validation"
	helper "github.com/salihkemaloglu/gignox-rr-beta-001/services"
)

func  SendEmailController(ctx context.Context, req *gigxRR.SendEmailRequest) (*gigxRR.SendEmailResponse, error) {
	userLang :="en"
	if headers, ok := metadata.FromIncomingContext(ctx); ok {
		userLang = headers["language"][0]
	}
	lang := helper.DetectLanguage(userLang)
	mailData := req.GetEmail();

	if valResp := val.SendMailFieldValidation(mailData.GetEmailAddress(),mailData.GetEmailType(),lang); valResp != "ok" {
		return nil,status.Errorf(
			codes.FailedPrecondition,
			fmt.Sprintf(valResp),
		)
	}

	verificationCode,verErr:=helper.GenerateVerificationCodeService()
	if verErr !=nil {
		verificationCode = "134584"
	}

	isOk:=false
	mailResp:=""
	if mailData.GetEmailType() == "register" {
		mailResp=helper.SendUserRegisterConfirmationMailService(mailData.GetEmailAddress(),verificationCode,userLang);
		if mailResp != "ok" {
			isOk=true
		}
	} else if mailData.GetEmailType() == "forgot" {
		mailResp=helper.SendUserForgotPasswordVerificationMailService(mailData.GetEmailAddress(),verificationCode,userLang);
		if mailResp != "ok" {
			isOk=true
		}
	} else {
		return nil,status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf(helper.Translate(lang,"Unknown_Email_type")),
		)
	} 

	return &gigxRR.SendEmailResponse{
		GeneralResponse:&gigxRR.GeneralResponse{
			Message:mailResp,
			IsEmailSuccess:isOk,
			IsOperationSuccess:true,
		},
	}, nil
}