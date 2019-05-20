package service

import (
	"fmt"
	"time"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/salihkemaloglu/gignox-rr-beta-001/proto"
	repo "github.com/salihkemaloglu/gignox-rr-beta-001/repositories"
	inter "github.com/salihkemaloglu/gignox-rr-beta-001/interfaces"
)
const (
    timeFormat = "2006-01-02 15:04:05"
)
func CheckVerificationTokenService(userTemporaryInformation_ *repo.UserTemporaryInformation,lang_ string) (*gigxRR.CheckVerificationTokenResponse, error) {

	var userTemporaryInformationOp inter.IUserTemporaryInformationRepository=userTemporaryInformation_
	if userTemporaryInformation_.EmailType == "register" {
		userTemporaryInformation,err := userTemporaryInformationOp.CheckRegisterVerificationToken()
		if err == nil  {
			return nil,status.Errorf(
				codes.NotFound,
				fmt.Sprintf(Translate(lang_,"register_verification_token_not_found")),
			)
		} else if userTemporaryInformation.IsTokenUsed {
			return nil,status.Errorf(
				codes.ResourceExhausted,
				fmt.Sprintf(Translate(lang_,"register_verification_token_used_before")),
			)
		} else if userTemporaryInformation.IsTokenExpired {
			return nil,status.Errorf(
				codes.DeadlineExceeded,
				fmt.Sprintf(Translate(lang_,"register_verification_token_expired")),
			)
		} else {
			timeNow := time.Now().UTC()
			createdTime, errTime := time.Parse(timeFormat, userTemporaryInformation.RegisterVerificationTokenCreateDate)
			if errTime != nil {
				return nil,status.Errorf(
					codes.Aborted,
					fmt.Sprintf(Translate(lang_,"register_verification_token_expired")),
				)
			}
			if minutes := timeNow.Sub(createdTime).Minutes(); minutes >= 120 {
				userTemporaryInformation.IsTokenUsed=false
				userTemporaryInformation.IsTokenExpired=true
				if updateErr := userTemporaryInformationOp.Update(); updateErr != nil {
					return nil,status.Errorf(
						codes.Aborted,
						fmt.Sprintf(Translate(lang_,"Register_Verification_Token_User_Temporary_Information_Update_Database_Error")),
					)
				}
				return nil,status.Errorf(
					codes.DeadlineExceeded,
					fmt.Sprintf(Translate(lang_,"register_verification_token_expired")),
				)
			} else {
				user := repo.User {
					Email: userTemporaryInformation_.Email,
				}
				var userOp inter.IUserRepository=user
				userGet,errUserGet := userOp.GetUserByEmail()
				if errUserGet != nil {
					return nil,status.Errorf(
						codes.NotFound,
						fmt.Sprintf(Translate(lang_,"register_verification_token_user_not_exist")),
					)
				}
				userGet.IsUserVerificated=true
				if updateErr := userOp.Update(); updateErr != nil {
					return nil,status.Errorf(
						codes.Aborted,
						fmt.Sprintf(Translate(lang_,"Register_Verification_Token_User_Update_Database_Error")),
					)
				}
				var userTemporaryInformationUpdateOp inter.IUserTemporaryInformationRepository=userTemporaryInformation				
				userTemporaryInformation.IsTokenUsed=true
				userTemporaryInformation.IsTokenExpired=true
				if updateErr := userTemporaryInformationUpdateOp.Update(); updateErr != nil {
					return nil,status.Errorf(
						codes.Aborted,
						fmt.Sprintf(Translate(lang_,"Register_Verification_Token_User_Temporary_Information_Update_Database_Error")),
					)
				}
				return &gigxRR.CheckVerificationTokenResponse {
					GeneralResponse:&gigxRR.GeneralResponse {
						IsOperationSuccess:true,
					},
				}, nil
			}
		}
		
	} else if userTemporaryInformation_.EmailType == "forgot" {
		userTemporaryInformation,err := userTemporaryInformationOp.CheckForgotPasswordVerificationToken()
		if err != nil  {
			return nil,status.Errorf(
				codes.NotFound,
				fmt.Sprintf(Translate(lang_,"forgot_password_verification_token_not_found")),
			)
		} else if userTemporaryInformation.IsTokenUsed {
			return nil,status.Errorf(
				codes.ResourceExhausted,
				fmt.Sprintf(Translate(lang_,"forgot_password_verification_token_used_before")),
			)
		} else if userTemporaryInformation.IsTokenExpired {
			return nil,status.Errorf(
				codes.DeadlineExceeded,
				fmt.Sprintf(Translate(lang_,"forgot_password_verification_token_expired")),
			)
		} else {
			timeNow := time.Now().UTC()
			createdTime, errTime := time.Parse(timeFormat, userTemporaryInformation.ForgotPasswordVerificationTokenCreateDate)
			if errTime != nil {
				return nil,status.Errorf(
					codes.Aborted,
					fmt.Sprintf(Translate(lang_,"forgot_password_verification_token_time_parse_error")),
				)
			}
			if minutes := timeNow.Sub(createdTime).Minutes(); minutes >= 120 {
				var userTemporaryInformationUpdateOp inter.IUserTemporaryInformationRepository=userTemporaryInformation		
				userTemporaryInformation.IsTokenUsed=false
				userTemporaryInformation.IsTokenExpired=true
				if updateErr := userTemporaryInformationUpdateOp.Update(); updateErr != nil {
					return nil,status.Errorf(
						codes.Aborted,
						fmt.Sprintf(Translate(lang_,"forgot_password_verification_token_user_temporary_information_update_database_error")+updateErr.Error()),
					)
				}
				return nil,status.Errorf(
					codes.DeadlineExceeded,
					fmt.Sprintf(Translate(lang_,"forgot_password_verification_token_expired")),
				)
			} else {
			
				return &gigxRR.CheckVerificationTokenResponse {
					GeneralResponse:&gigxRR.GeneralResponse {
						IsOperationSuccess:true,
						Message:userTemporaryInformation.Email,
					},
				}, nil
			}
		}
	} else {
		return nil,status.Errorf(
			codes.Unknown,
			fmt.Sprintf(Translate(lang_,"unknown_email_type")),
		)
	}
}