package controller

import (
	"fmt"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/metadata"
	"github.com/salihkemaloglu/gignox-rr-beta-001/proto"
	repo "github.com/salihkemaloglu/gignox-rr-beta-001/repositories"
	inter "github.com/salihkemaloglu/gignox-rr-beta-001/interfaces"
	val "github.com/salihkemaloglu/gignox-rr-beta-001/validations"
	helper "github.com/salihkemaloglu/gignox-rr-beta-001/services"
)

func ResetUserPasswordController(ctx_ context.Context, req_ *gigxRR.ResetUserPasswordRequest) (*gigxRR.ResetUserPasswordResponse, error) {
	userLang :="en"
	if headers, ok := metadata.FromIncomingContext(ctx_); ok {
		userLang = headers["language"][0]
	}
	lang := helper.DetectLanguage(userLang)
	generalRequest := req_.GetGeneralRequest();
	userTemporaryInformation:= repo.UserTemporaryInformation {
		Email: generalRequest.GetEmailAddress(),
		RegisterVerificationToken: generalRequest.GetRegisterVerificationToken(),
		ForgotPasswordVerificationToken: generalRequest.GetForgotPasswordVerificationToken(),
		EmailType:generalRequest.GetEmailType(),
	}
	if valResp := val.ResetUserPasswordValidation(&userTemporaryInformation,generalRequest.GetPassword(),generalRequest.GetPasswordConfirm(),lang); valResp != "ok" {
		return nil,status.Errorf(
			codes.FailedPrecondition,
			fmt.Sprintf(valResp),
		)
	}

	resp,err:=helper.CheckVerificationTokenService(&userTemporaryInformation,lang)
	if err != nil {
		return nil,err
	}

    if resp.GetGeneralResponse().GetIsOperationSuccess() {
		user := repo.User {
			Email:userTemporaryInformation.Email,
		}
		var userOp inter.IUserRepository = &user
		userResp,userErr := userOp.GetUserByEmail()
		if userErr != nil {
			return nil,status.Errorf(
				codes.Aborted,
				fmt.Sprintf(helper.Translate(lang,"Forgot_Password_Invalid_User")+userErr.Error()),
			)
		}

		userResp.Password=generalRequest.GetPassword()
		var userUpdateOp inter.IUserRepository = userResp
		userErr=userUpdateOp.UpdateUserPassword()
		if userErr != nil {
			return nil,status.Errorf(
				codes.Aborted,
				fmt.Sprintf(helper.Translate(lang,"Forgot_Password_Invalid_User")+userErr.Error()),
			)
		}
		userTemporaryInformation.IsTokenUsed=true
		var userTemporaryInformationOp inter.IUserTemporaryInformationRepository=userTemporaryInformation
		if updateErr := userTemporaryInformationOp.UpdateByEmail(); updateErr != nil {
			return nil,status.Errorf(
				codes.Aborted,
				fmt.Sprintf(helper.Translate(lang,"Forgot_Password_Reset_New_Password_Database_Update_Error")+updateErr.Error()),
			)
		}
		return &gigxRR.ResetUserPasswordResponse {
			GeneralResponse:&gigxRR.GeneralResponse {
				IsOperationSuccess:true,
			},
		}, nil

	}
  
	return nil,status.Errorf(
		codes.Aborted,
		fmt.Sprintf(helper.Translate(lang,"Unknown_Service_Error")),
	)

}