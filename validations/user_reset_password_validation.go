package validation

import (
	"strings"

	repo "github.com/salihkemaloglu/gignoxrr-beta-001/repositories"
	helper "github.com/salihkemaloglu/gignoxrr-beta-001/services"
	zxcvbn "github.com/nbutton23/zxcvbn-go"
)

func ResetUserPasswordValidation(userTemporaryInformation_ *repo.UserTemporaryInformation,password_ string,passwordConfirm_ string,lang_ string) string {

	if strings.TrimSpace(userTemporaryInformation_.Email) == "" {
		return helper.Translate(lang_,"email_field_empty_validation")
	} else if strings.TrimSpace(userTemporaryInformation_.ForgotPasswordVerificationToken) == "" {
		return helper.Translate(lang_,"forgot_password_verification_token_field_empty_validation")
	} else if strings.TrimSpace(password_) == "" {
		return helper.Translate(lang_,"password_field_empty_validation")
	}  else if strings.TrimSpace(passwordConfirm_) == "" {
		return helper.Translate(lang_,"forgot_password_verification_token_field_empty_validation")
	} else if password_ != passwordConfirm_ {
		return helper.Translate(lang_,"password_and_confirm_not_match")
	}else if zxcvbn.PasswordStrength(password_,nil ).Score < 1 {
		return helper.Translate(lang_,"password_reset_strenght_validation_information")
	} else {
		return "ok"
	}
	
}