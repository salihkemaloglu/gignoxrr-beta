package validations

import (
	"strings"

	helper "github.com/salihkemaloglu/gignoxrr-beta-001/helpers"
	repo "github.com/salihkemaloglu/gignoxrr-beta-001/repositories"
)

//SendMailFieldValidation ...
func SendMailFieldValidation(email string, username string, mailType string, lang string) string {

	if strings.TrimSpace(email) == "" {
		return helper.Translate(lang, "email_field_empty_validation")
	} else if strings.TrimSpace(mailType) == "" {
		return helper.Translate(lang, "email_type_field_empty_validation")
	} else if mailType == "register" {
		if strings.TrimSpace(username) == "" {
			return helper.Translate(lang, "username_field_empty_validation")
		}
		return "ok"
	} else {
		return "ok"
	}
}

//CheckVerificationTokenValidation ...
func CheckVerificationTokenValidation(userTemporaryInformation *repo.UserTemporaryInformation, lang string) string {

	if strings.TrimSpace(userTemporaryInformation.EmailType) == "" {
		return helper.Translate(lang, "email_type_field_empty_validation")
	} else if userTemporaryInformation.EmailType == "forgot" {
		if strings.TrimSpace(userTemporaryInformation.ForgotPasswordVerificationToken) == "" {
			return helper.Translate(lang, "forgot_password_verification_token_field_empty_validation")
		}
		return "ok"

	} else if userTemporaryInformation.EmailType == "register" {
		if strings.TrimSpace(userTemporaryInformation.RegisterVerificationToken) == "" {
			return helper.Translate(lang, "register_verification_token_field_empty_validation")
		}
		return "ok"

	} else {
		return "ok"
	}
}
