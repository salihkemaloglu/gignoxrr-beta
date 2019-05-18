package validation

import (
	"strings"

	repo "github.com/salihkemaloglu/gignox-rr-beta-001/repositories"
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
func CheckVerificationTokenValidation(userTemporaryInformation_ *repo.UserTemporaryInformation,lang_ string) string {

	if strings.TrimSpace(userTemporaryInformation_.EmailType) == "" {
		return helper.Translate(lang_,"Email_Type_Field_Empty_Validation")
	} else if userTemporaryInformation_.EmailType == "forgot" {
		if strings.TrimSpace(userTemporaryInformation_.ForgotPasswordVerificationToken) == "" {
			return helper.Translate(lang_,"Forgot_Password_Verification_Token_Field_Empty_Validation")
		} else {
			return "ok"
		}
	} else if userTemporaryInformation_.EmailType == "register" {
		if strings.TrimSpace(userTemporaryInformation_.RegisterVerificationToken) == "" {
			return helper.Translate(lang_,"Register_Verification_Token_Field_Empty_Validation")
		} else {
			return "ok"
		}
	} else {
		return "ok"
	}
}