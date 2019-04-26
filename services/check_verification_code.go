package service

import (
	"time"
	"fmt"
	db "github.com/salihkemaloglu/gignox-rr-beta-001/mongodb"
	repo "github.com/salihkemaloglu/gignox-rr-beta-001/repository"
)
const (
    timeFormat = "2006-01-02 15:04:05"
)
func CheckVerificationCodeService(userTemporaryInformation_ *db.UserTemporaryInformation,mailType_ string,lang_ string) (string,error) {

	var _userTemporaryInformationOp repo.UserTemporaryInformationRepository=userTemporaryInformation_
	if mailType_ == "register" {
		_userTemporaryInformation,_err := _userTemporaryInformationOp.CheckRegisterVerificationCode()
		if _err == nil  {
			return fmt.Sprintf(Translate(lang_,"Register_Verification_Code_Not_Found")),_err
		} else if _userTemporaryInformation.IsCodeUsed {
			return fmt.Sprintf(Translate(lang_,"Register_Verification_Code_Used_Before")),nil
		} else if _userTemporaryInformation.IsCodeExpired {
			return fmt.Sprintf(Translate(lang_,"Register_Verification_Code_Expired")),nil
		} else {
			_timeNow := time.Now().UTC()
			_createdTime, _errTime := time.Parse(timeFormat, _userTemporaryInformation.RegisterVerificationCodeCreateDate)
			if _errTime != nil {
				return fmt.Sprintf(Translate(lang_,"Register_Verification_Code_Time_Parse_Error")),nil
			}
			if _minutes := _timeNow.Sub(_createdTime).Minutes(); _minutes >= 120 {
				_userTemporaryInformation.IsCodeUsed=true
				_userTemporaryInformation.IsCodeExpired=true
				if _updateErr := _userTemporaryInformationOp.Update(); _updateErr != nil {
					return fmt.Sprintf(Translate(lang_,"Register_Verification_Code_User_Temporary_Information_Update_Error")),_updateErr
				}
				return fmt.Sprintf(Translate(lang_,"Register_Verification_Code_Expired")),nil
			} else {
				_user := db.User {
					Email: userTemporaryInformation_.Email,
				}
				var _userOp repo.UserRepository=_user
				_userGet,_errUserGet := _userOp.GetUserByEmail()
				if _errUserGet != nil {
					return fmt.Sprintf(Translate(lang_,"Register_Verification_Code_User_Not_Exist")),_errUserGet
				}
				_userGet.IsUserVerificated=true
				if _updateErr := _userOp.Update(); _updateErr != nil {
					return fmt.Sprintf(Translate(lang_,"Register_Verification_Code_User_Update_Error")),_updateErr
				}
				_userTemporaryInformation.IsCodeUsed=true
				_userTemporaryInformation.IsCodeExpired=true
				if _updateErr := _userTemporaryInformationOp.Update(); _updateErr != nil {
					return fmt.Sprintf(Translate(lang_,"Register_Verification_Code_User_Temporary_Information_Update_Error")),_updateErr
				}
				return "ok",nil
			}
		}
		
	} else if mailType_ == "forgot" {
		_userTemporaryInformation,_err := _userTemporaryInformationOp.CheckForgotPasswordVerificationCode()
		if _err == nil  {
			return fmt.Sprintf(Translate(lang_,"Forgot_Password_Verification_Code_Not_Found")),_err
		} else if _userTemporaryInformation.IsCodeUsed {
			return fmt.Sprintf(Translate(lang_,"Register_Verification_Code_Used_Before")),nil
		} else if _userTemporaryInformation.IsCodeExpired {
			return fmt.Sprintf(Translate(lang_,"Register_Verification_Code_Expired")),nil
		} else {
			_timeNow := time.Now().UTC()
			_createdTime, _errTime := time.Parse(timeFormat, _userTemporaryInformation.ForgotPasswordVerificationCode)
			if _errTime != nil {
				return fmt.Sprintf(Translate(lang_,"Register_Verification_Code_Time_Parse_Error")),nil
			}
			if _minutes := _timeNow.Sub(_createdTime).Minutes(); _minutes >= 120 {
				_userTemporaryInformation.IsCodeUsed=true
				_userTemporaryInformation.IsCodeExpired=true
				if _updateErr := _userTemporaryInformationOp.Update(); _updateErr != nil {
					return fmt.Sprintf(Translate(lang_,"Register_Verification_Code_User_Temporary_Information_Update_Error")),_updateErr
				}
				return fmt.Sprintf(Translate(lang_,"Register_Verification_Code_Expired")),nil
			} else {
				_userTemporaryInformation.IsCodeUsed=true
				_userTemporaryInformation.IsCodeExpired=true
				if _updateErr := _userTemporaryInformationOp.Update(); _updateErr != nil {
					return fmt.Sprintf(Translate(lang_,"Register_Verification_Code_User_Temporary_Information_Update_Error")),_updateErr
				}
				return "ok",nil
			}
		}
	} else {
		return 	fmt.Sprintf(Translate(lang_,"Unknown_Email_type")),nil
	}
}