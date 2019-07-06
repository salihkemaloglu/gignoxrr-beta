package validations

import (
	"strings"

	zxcvbn "github.com/nbutton23/zxcvbn-go"
	helper "github.com/salihkemaloglu/gignoxrr-beta-001/helpers"
	repo "github.com/salihkemaloglu/gignoxrr-beta-001/repositories"
)

//ResetUserPasswordValidation ...
func ResetUserPasswordFieldValidation(userTemporaryInformation *repo.UserTemporaryInformation, password string, passwordConfirm string, lang string) string {

	if strings.TrimSpace(userTemporaryInformation.Email) == "" {
		return helper.Translate(lang, "email_field_empty_validation")
	} else if strings.TrimSpace(userTemporaryInformation.ForgotPasswordVerificationToken) == "" {
		return helper.Translate(lang, "forgot_password_verification_token_field_empty_validation")
	} else if strings.TrimSpace(password) == "" {
		return helper.Translate(lang, "password_field_empty_validation")
	} else if strings.TrimSpace(passwordConfirm) == "" {
		return helper.Translate(lang, "forgot_password_verification_token_field_empty_validation")
	} else if password != passwordConfirm {
		return helper.Translate(lang, "password_and_confirm_not_match")
	} else if zxcvbn.PasswordStrength(password, nil).Score < 1 {
		return helper.Translate(lang, "password_reset_strenght_validation_information")
	} else {
		return "ok"
	}

}
