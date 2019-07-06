package services

import (
	"context"
	"fmt"

	helper "github.com/salihkemaloglu/gignoxrr-beta-001/helpers"
	gigxRR "github.com/salihkemaloglu/gignoxrr-beta-001/proto"
	repo "github.com/salihkemaloglu/gignoxrr-beta-001/repositories"
	val "github.com/salihkemaloglu/gignoxrr-beta-001/validations"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

//CheckVerificationLinkService ...
func CheckVerificationLinkService(ctx context.Context, req *gigxRR.CheckVerificationLinkRequest) (*gigxRR.CheckVerificationLinkResponse, error) {
	userLang := "en"
	if headers, ok := metadata.FromIncomingContext(ctx); ok {
		if headers["languagecode"] != nil {
			userLang = headers["languagecode"][0]
		}
	}
	lang := helper.DetectLanguage(userLang)
	emailData := req.GetGeneralRequest()
	userTemporaryInformation := repo.UserTemporaryInformation{
		Email:                           emailData.GetEmailAddress(),
		RegisterVerificationToken:       emailData.GetRegisterVerificationToken(),
		ForgotPasswordVerificationToken: emailData.GetForgotPasswordVerificationToken(),
		EmailType:                       emailData.GetEmailType(),
	}
	if valResp := val.CheckVerificationTokenValidation(&userTemporaryInformation, lang); valResp != "ok" {
		return nil, status.Errorf(
			codes.FailedPrecondition,
			fmt.Sprintf(valResp),
		)
	}

	return helper.CheckVerificationLinkService(&userTemporaryInformation, lang)

}
