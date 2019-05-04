package validation

import (
	"strings"

	str "github.com/salihkemaloglu/gignox-rr-beta-001/repositories"
	helper "github.com/salihkemaloglu/gignox-rr-beta-001/services"
)

func SendMailFieldValidation(email_ string,mailType_ string,lang_ string) string{

	if strings.TrimSpace(email_) == "" {
		return helper.Translate(lang_,"Email_Field_Empty_Validation")
	} else if strings.TrimSpace(mailType_) == "" {
		return helper.Translate(lang_,"Email_Type_Field_Empty_Validation")
	} else {
		return "ok"
	}
}
func CheckVerificationCodeValidation(userTemporaryInformation_ *str.UserTemporaryInformation,mailType_ string,lang_ string) string {

	if strings.TrimSpace(userTemporaryInformation_.Email) == "" {
		return helper.Translate(lang_,"Email_Field_Empty_Validation")
	} else if strings.TrimSpace(mailType_) == "" {
		return helper.Translate(lang_,"Email_Type_Field_Empty_Validation")
	} else if mailType_ == "forgot" {
		if strings.TrimSpace(userTemporaryInformation_.ForgotPasswordVerificationCode) == "" {
			return helper.Translate(lang_,"Forgot_Password_Verification_Code_Field_Empty_Validation")
		} else {
			return "ok"
		}
	} else if mailType_ == "register" {
		if strings.TrimSpace(userTemporaryInformation_.RegisterVerificationCode) == "" {
			return helper.Translate(lang_,"Register_Verification_Code_Field_Empty_Validation")
		} else {
			return "ok"
		}
	} else {
		return "ok"
	}
}