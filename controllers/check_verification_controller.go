package controller

import (
	"fmt"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/metadata"
	"github.com/salihkemaloglu/gignox-rr-beta-001/proto"
	repo "github.com/salihkemaloglu/gignox-rr-beta-001/repositories"
	val "github.com/salihkemaloglu/gignox-rr-beta-001/validations"
	helper "github.com/salihkemaloglu/gignox-rr-beta-001/services"
)

func CheckVerificationTokenController(ctx_ context.Context, req_ *gigxRR.CheckVerificationTokenRequest) (*gigxRR.CheckVerificationTokenResponse, error) {
	userLang :="en"
	if headers, ok := metadata.FromIncomingContext(ctx_); ok {
		userLang = headers["language"][0]
	}
	lang := helper.DetectLanguage(userLang)
	emailData := req_.GetGeneralRequest();
	userTemporaryInformation:=repo.UserTemporaryInformation {
		Email: emailData.GetEmailAddress(),
		RegisterVerificationToken: emailData.GetRegisterVerificationToken(),
		ForgotPasswordVerificationToken: emailData.GetForgotPasswordVerificationToken(),
		EmailType:emailData.GetEmailType(),
	}
	if valResp := val.CheckVerificationTokenValidation(&userTemporaryInformation,lang); valResp != "ok" {
		return nil,status.Errorf(
			codes.FailedPrecondition,
			fmt.Sprintf(valResp),
		)
	}

	return helper.CheckVerificationTokenService(&userTemporaryInformation,lang)

}