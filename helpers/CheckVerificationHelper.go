package helper

import (
	"fmt"
	"time"

	inter "github.com/salihkemaloglu/gignoxrr-beta-001/interfaces"
	gigxRR "github.com/salihkemaloglu/gignoxrr-beta-001/proto"
	repo "github.com/salihkemaloglu/gignoxrr-beta-001/repositories"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	timeFormat = "2006-01-02 15:04:05"
)

//CheckVerificationLinkService ...
func CheckVerificationLinkService(userTemporaryInformation *repo.UserTemporaryInformation, lang string) (*gigxRR.CheckVerificationLinkResponse, error) {

	var userTemporaryInformationOp inter.IUserTemporaryInformationRepository = userTemporaryInformation
	if userTemporaryInformation.EmailType == "register" {
		userTemporaryInformation, err := userTemporaryInformationOp.CheckRegisterVerificationToken()
		if err == nil {
			return nil, status.Errorf(
				codes.NotFound,
				fmt.Sprintf(Translate(lang, "register_verification_token_not_found")),
			)
		} else if userTemporaryInformation.IsTokenUsed {
			return nil, status.Errorf(
				codes.ResourceExhausted,
				fmt.Sprintf(Translate(lang, "register_verification_token_used_before")),
			)
		} else if userTemporaryInformation.IsTokenExpired {
			return nil, status.Errorf(
				codes.DeadlineExceeded,
				fmt.Sprintf(Translate(lang, "register_verification_token_expired")),
			)
		} else {
			user := repo.User{
				Email: userTemporaryInformation.Email,
			}
			var userOp inter.IUserRepository = user
			userGet, errUserGet := userOp.GetUserByEmail()
			if errUserGet != nil {
				return nil, status.Errorf(
					codes.NotFound,
					fmt.Sprintf(Translate(lang, "register_verification_token_user_not_exist")),
				)
			}
			userGet.IsAccountConfirm = true
			var userUpdateOp inter.IUserRepository = userGet
			if updateErr := userUpdateOp.Update(); updateErr != nil {
				return nil, status.Errorf(
					codes.Aborted,
					fmt.Sprintf(Translate(lang, "Register_Verification_Token_User_Update_Database_Error")),
				)
			}
			var userTemporaryInformationUpdateOp inter.IUserTemporaryInformationRepository = userTemporaryInformation
			userTemporaryInformation.IsTokenUsed = true
			userTemporaryInformation.IsTokenExpired = true
			if updateErr := userTemporaryInformationUpdateOp.Update(); updateErr != nil {
				return nil, status.Errorf(
					codes.Aborted,
					fmt.Sprintf(Translate(lang, "Register_Verification_Token_User_Temporary_Information_Update_Database_Error")),
				)
			}
			token, tokenErr := CreateTokenEndpointService(user)
			isTokenSuccess := true
			if tokenErr != nil {
				isTokenSuccess = false
				token = tokenErr.Error()
			}
			return &gigxRR.CheckVerificationLinkResponse{
				GeneralResponse: &gigxRR.GeneralResponse{
					Message:            "verify user account",
					Token:              token,
					IsTokenSuccess:     isTokenSuccess,
					IsOperationSuccess: true,
				},
			}, nil
		}

	} else if userTemporaryInformation.EmailType == "forgot" {
		userTemporaryInformation, err := userTemporaryInformationOp.CheckForgotPasswordVerificationToken()
		if err != nil {
			return nil, status.Errorf(
				codes.NotFound,
				fmt.Sprintf(Translate(lang, "forgot_password_verification_token_not_found")),
			)
		} else if userTemporaryInformation.IsTokenUsed {
			return nil, status.Errorf(
				codes.ResourceExhausted,
				fmt.Sprintf(Translate(lang, "forgot_password_verification_token_used_before")),
			)
		} else if userTemporaryInformation.IsTokenExpired {
			return nil, status.Errorf(
				codes.DeadlineExceeded,
				fmt.Sprintf(Translate(lang, "forgot_password_verification_token_expired")),
			)
		} else {
			timeNow := time.Now().UTC()
			createdTime, errTime := time.Parse(timeFormat, userTemporaryInformation.ForgotPasswordVerificationTokenCreateDate)
			if errTime != nil {
				return nil, status.Errorf(
					codes.Aborted,
					fmt.Sprintf(Translate(lang, "forgot_password_verification_token_time_parse_error")),
				)
			}
			if minutes := timeNow.Sub(createdTime).Minutes(); minutes >= 120 {
				var userTemporaryInformationUpdateOp inter.IUserTemporaryInformationRepository = userTemporaryInformation
				userTemporaryInformation.IsTokenUsed = false
				userTemporaryInformation.IsTokenExpired = true
				if updateErr := userTemporaryInformationUpdateOp.Update(); updateErr != nil {
					return nil, status.Errorf(
						codes.Aborted,
						fmt.Sprintf(Translate(lang, "forgot_password_verification_token_user_temporary_information_update_database_error")+updateErr.Error()),
					)
				}
				return nil, status.Errorf(
					codes.DeadlineExceeded,
					fmt.Sprintf(Translate(lang, "forgot_password_verification_token_expired")),
				)
			}
			return &gigxRR.CheckVerificationLinkResponse{
				GeneralResponse: &gigxRR.GeneralResponse{
					IsOperationSuccess: true,
					Message:            userTemporaryInformation.Email,
				},
			}, nil

		}
	} else {
		return nil, status.Errorf(
			codes.Unknown,
			fmt.Sprintf(Translate(lang, "unknown_email_type")),
		)
	}
}
