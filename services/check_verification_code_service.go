package service

import (
	"time"
	"fmt"
	repo "github.com/salihkemaloglu/gignox-rr-beta-001/repositories"
	inter "github.com/salihkemaloglu/gignox-rr-beta-001/interfaces"
)
const (
    timeFormat = "2006-01-02 15:04:05"
)
func CheckVerificationCodeService(userTemporaryInformation_ *repo.UserTemporaryInformation,mailType_ string,lang_ string) (string,error) {

	var userTemporaryInformationOp inter.IUserTemporaryInformationRepository=userTemporaryInformation_
	if mailType_ == "register" {
		userTemporaryInformation,err := userTemporaryInformationOp.CheckRegisterVerificationCode()
		if err == nil  {
			return fmt.Sprintf(Translate(lang_,"Register_Verification_Code_Not_Found")),err
		} else if userTemporaryInformation.IsCodeUsed {
			return fmt.Sprintf(Translate(lang_,"Register_Verification_Code_Used_Before")),nil
		} else if userTemporaryInformation.IsCodeExpired {
			return fmt.Sprintf(Translate(lang_,"Register_Verification_Code_Expired")),nil
		} else {
			timeNow := time.Now().UTC()
			createdTime, errTime := time.Parse(timeFormat, userTemporaryInformation.RegisterVerificationCodeCreateDate)
			if errTime != nil {
				return fmt.Sprintf(Translate(lang_,"Register_Verification_Code_Time_Parse_Error")),nil
			}
			if minutes := timeNow.Sub(createdTime).Minutes(); minutes >= 120 {
				userTemporaryInformation.IsCodeUsed=true
				userTemporaryInformation.IsCodeExpired=true
				if updateErr := userTemporaryInformationOp.Update(); updateErr != nil {
					return fmt.Sprintf(Translate(lang_,"Register_Verification_Code_User_Temporary_Information_Update_Error")),updateErr
				}
				return fmt.Sprintf(Translate(lang_,"Register_Verification_Code_Expired")),nil
			} else {
				user := repo.User {
					Email: userTemporaryInformation_.Email,
				}
				var userOp inter.IUserRepository=user
				userGet,errUserGet := userOp.GetUserByEmail()
				if errUserGet != nil {
					return fmt.Sprintf(Translate(lang_,"Register_Verification_Code_User_Not_Exist")),errUserGet
				}
				userGet.IsUserVerificated=true
				if updateErr := userOp.Update(); updateErr != nil {
					return fmt.Sprintf(Translate(lang_,"Register_Verification_Code_User_Update_Error")),updateErr
				}
				userTemporaryInformation.IsCodeUsed=true
				userTemporaryInformation.IsCodeExpired=true
				if updateErr := userTemporaryInformationOp.Update(); updateErr != nil {
					return fmt.Sprintf(Translate(lang_,"Register_Verification_Code_User_Temporary_Information_Update_Error")),updateErr
				}
				return "ok",nil
			}
		}
		
	} else if mailType_ == "forgot" {
		userTemporaryInformation,err := userTemporaryInformationOp.CheckForgotPasswordVerificationCode()
		if err == nil  {
			return fmt.Sprintf(Translate(lang_,"Forgot_Password_Verification_Code_Not_Found")),err
		} else if userTemporaryInformation.IsCodeUsed {
			return fmt.Sprintf(Translate(lang_,"Forgot_Password_Verification_Code_Used_Before")),nil
		} else if userTemporaryInformation.IsCodeExpired {
			return fmt.Sprintf(Translate(lang_,"Forgot_Password_Verification_Code_Expired")),nil
		} else {
			timeNow := time.Now().UTC()
			createdTime, errTime := time.Parse(timeFormat, userTemporaryInformation.ForgotPasswordVerificationCode)
			if errTime != nil {
				return fmt.Sprintf(Translate(lang_,"Forgot_Password_Verification_Code_Time_Parse_Error")),nil
			}
			if minutes := timeNow.Sub(createdTime).Minutes(); minutes >= 120 {
				userTemporaryInformation.IsCodeUsed=true
				userTemporaryInformation.IsCodeExpired=true
				if updateErr := userTemporaryInformationOp.Update(); updateErr != nil {
					return fmt.Sprintf(Translate(lang_,"Forgot_Password_Verification_Code_User_Temporary_Information_Update_Error")),updateErr
				}
				return fmt.Sprintf(Translate(lang_,"Forgot_Password_Verification_Code_Expired")),nil
			} else {
				userTemporaryInformation.IsCodeUsed=true
				userTemporaryInformation.IsCodeExpired=true
				if updateErr := userTemporaryInformationOp.Update(); updateErr != nil {
					return fmt.Sprintf(Translate(lang_,"Forgot_Password_Verification_Code_User_Temporary_Information_Update_Error")),updateErr
				}
				return "ok",nil
			}
		}
	} else {
		return 	fmt.Sprintf(Translate(lang_,"Unknown_Email_type")),nil
	}
}