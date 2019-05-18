package validation

import (
	"strings"

	repo "github.com/salihkemaloglu/gignox-rr-beta-001/repositories"
	helper "github.com/salihkemaloglu/gignox-rr-beta-001/services"
)

func ResetUserPasswordValidation(userTemporaryInformation_ *repo.UserTemporaryInformation,password_ string,passwordConfirm_ string,lang_ string) string {

	if strings.TrimSpace(userTemporaryInformation_.Email) == "" {
		return helper.Translate(lang_,"Email_Field_Empty_Validation")
	} else if strings.TrimSpace(userTemporaryInformation_.ForgotPasswordVerificationToken) == "" {
		return helper.Translate(lang_,"Forgot_Password_Verification_Token_Field_Empty_Validation")
	} else if strings.TrimSpace(password_) == "" {
		return helper.Translate(lang_,"Password_Field_Empty_Validation")
	}  else if strings.TrimSpace(passwordConfirm_) == "" {
		return helper.Translate(lang_,"Forgot_Password_Verification_Token_Field_Empty_Validation")
	} else if password_ != passwordConfirm_ {
		return helper.Translate(lang_,"Password_And_Confirm_Not_Match")
	} else {
		return "ok"
	}
	
}