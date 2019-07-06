package services

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

//SendEmailService ...
func SendEmailService(ctx context.Context, req *gigxRR.SendEmailRequest) (*gigxRR.SendEmailResponse, error) {
	userLang := "en"
	if headers, ok := metadata.FromIncomingContext(ctx); ok {
		if headers["languagecode"] != nil {
			userLang = headers["languagecode"][0]
		}
	}
	lang := helper.DetectLanguage(userLang)

	mailData := req.GetGeneralRequest()
	user := repo.User{
		Email: mailData.GetEmailAddress(),
	}

	if valResp := val.SendMailFieldValidation(mailData.GetEmailAddress(), mailData.GetUsername(), mailData.GetEmailType(), lang); valResp != "ok" {
		return nil, status.Errorf(
			codes.FailedPrecondition,
			fmt.Sprintf(valResp),
		)
	}

	var userOp inter.IUserRepository = user
	if err := userOp.CheckUser(); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf(helper.Translate(lang, "user_account_not_exist_account")+user.Email),
		)
	}

	verificationToken, verErr := helper.GenerateRandomStringURLService(128)
	if verErr != nil {
		return nil, status.Errorf(
			codes.Aborted,
			fmt.Sprintf(helper.Translate(lang, "generate_password_verification_token_error")+verErr.Error()),
		)
	}

	if mailData.GetEmailType() == "register" {
		return helper.SendUserRegisterConfirmationMailService(mailData.GetEmailAddress(), mailData.GetUsername(), mailData.GetEmailType(), verificationToken, userLang)

	} else if mailData.GetEmailType() == "forgot" {
		return helper.SendUserForgotPasswordVerificationMailService(mailData.GetEmailAddress(), mailData.GetEmailType(), verificationToken, userLang)

	} else {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf(helper.Translate(lang, "unknown_email_type")),
		)
	}

}
