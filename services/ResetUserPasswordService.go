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

//ResetUserPasswordService ...
func ResetUserPasswordService(ctx context.Context, req *gigxRR.ResetUserPasswordRequest) (*gigxRR.ResetUserPasswordResponse, error) {
	userLang := "en"
	if headers, ok := metadata.FromIncomingContext(ctx); ok {
		if headers["languagecode"] != nil {
			userLang = headers["languagecode"][0]
		}
	}
	lang := helper.DetectLanguage(userLang)
	generalRequest := req.GetGeneralRequest()
	userTemporaryInformation := repo.UserTemporaryInformation{
		Email:                           generalRequest.GetEmailAddress(),
		RegisterVerificationToken:       generalRequest.GetRegisterVerificationToken(),
		ForgotPasswordVerificationToken: generalRequest.GetForgotPasswordVerificationToken(),
		EmailType:                       generalRequest.GetEmailType(),
	}
	if valResp := val.ResetUserPasswordFieldValidation(&userTemporaryInformation, generalRequest.GetPassword(), generalRequest.GetPasswordConfirm(), lang); valResp != "ok" {
		return nil, status.Errorf(
			codes.FailedPrecondition,
			fmt.Sprintf(valResp),
		)
	}

	resp, err := helper.CheckVerificationLinkService(&userTemporaryInformation, lang)
	if err != nil {
		return nil, err
	}

	if resp.GetGeneralResponse().GetIsOperationSuccess() {
		user := repo.User{
			Email: userTemporaryInformation.Email,
		}
		var userOp inter.IUserRepository = &user
		userResp, userErr := userOp.GetUserByEmail()
		if userErr != nil {
			return nil, status.Errorf(
				codes.InvalidArgument,
				fmt.Sprintf(helper.Translate(lang, "forgot_password_invalid_user")+userErr.Error()),
			)
		}

		userResp.Password = helper.EncryptePassword(generalRequest.GetPassword())
		var userUpdateOp inter.IUserRepository = userResp
		userErr = userUpdateOp.UpdateUserPassword()
		if userErr != nil {
			return nil, status.Errorf(
				codes.Aborted,
				fmt.Sprintf(helper.Translate(lang, "forgot_password_reset_new_password_database_update_error")+userErr.Error()),
			)
		}
		userTemporaryInformation.IsTokenUsed = true
		var userTemporaryInformationOp inter.IUserTemporaryInformationRepository = userTemporaryInformation
		if updateErr := userTemporaryInformationOp.UpdateByEmail(); updateErr != nil {
			return nil, status.Errorf(
				codes.Aborted,
				fmt.Sprintf(helper.Translate(lang, "forgot_password_reset_new_password_database_update_error")+updateErr.Error()),
			)
		}
		return &gigxRR.ResetUserPasswordResponse{
			GeneralResponse: &gigxRR.GeneralResponse{
				IsOperationSuccess: true,
			},
		}, nil

	}

	return nil, status.Errorf(
		codes.Aborted,
		fmt.Sprintf(helper.Translate(lang, "unknown_service_error")),
	)

}
