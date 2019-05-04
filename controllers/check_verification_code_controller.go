package controller

import (
	"fmt"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/metadata"
	"github.com/salihkemaloglu/gignox-rr-beta-001/proto"
	str "github.com/salihkemaloglu/gignox-rr-beta-001/repositories"
	val "github.com/salihkemaloglu/gignox-rr-beta-001/validations"
	helper "github.com/salihkemaloglu/gignox-rr-beta-001/services"
)

func CheckVerificationCodeController(ctx_ context.Context, req_ *gigxRR.CheckVerificationCodeRequest) (*gigxRR.CheckVerificationCodeResponse, error) {
	userLang :="en"
	if headers, ok := metadata.FromIncomingContext(ctx_); ok {
		userLang = headers["language"][0]
	}
	lang := helper.DetectLanguage(userLang)
	emailData := req_.GetEmail();
	userTemporaryInformation:=str.UserTemporaryInformation {
		Email: emailData.GetEmailAddress(),
		RegisterVerificationCode: emailData.GetRegisterVerificationCode(),
		ForgotPasswordVerificationCode: emailData.GetForgotPasswordVerificationCode(),
	}
	if valResp := val.CheckVerificationCodeValidation(&userTemporaryInformation,emailData.GetEmailType(),lang); valResp != "ok" {
		return nil,status.Errorf(
			codes.FailedPrecondition,
			fmt.Sprintf(valResp),
		)
	}

	verificationCode,verErr:=helper.CheckVerificationCodeService(&userTemporaryInformation,emailData.GetEmailType(),lang)
	if verErr !=nil {
		verificationCode = "134584"
	}
	isOk:=false
	emailResp:=helper.SendUserRegisterConfirmationMailService(emailData.GetEmailAddress(),verificationCode,userLang);
	if emailResp != "ok" {
		isOk=true
	}

	return &gigxRR.CheckVerificationCodeResponse{
		GeneralResponse:&gigxRR.GeneralResponse{
			Message:emailResp,
			IsEmailSuccess:isOk,
			IsOperationSuccess:true,
		},
	}, nil
}