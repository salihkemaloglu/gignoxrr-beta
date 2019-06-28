package validation

import (
	"strings"

	repo "github.com/salihkemaloglu/gignoxrr-beta-001/repositories"
	helper "github.com/salihkemaloglu/gignoxrr-beta-001/services"
)

func SendMailFieldValidation(email_ string, username_ string, mailType_ string, lang_ string) string {

	if strings.TrimSpace(email_) == "" {
		return helper.Translate(lang_, "email_field_empty_validation")
	} else if strings.TrimSpace(mailType_) == "" {
		return helper.Translate(lang_, "email_type_field_empty_validation")
	} else if mailType_ == "register" {
		if strings.TrimSpace(username_) == "" {
			return helper.Translate(lang_, "username_field_empty_validation")
		} else {
			return "ok"
		}
	} else {
		return "ok"
	}
}
func CheckVerificationTokenValidation(userTemporaryInformation_ *repo.UserTemporaryInformation, lang_ string) string {

	if strings.TrimSpace(userTemporaryInformation_.EmailType) == "" {
		return helper.Translate(lang_, "email_type_field_empty_validation")
	} else if userTemporaryInformation_.EmailType == "forgot" {
		if strings.TrimSpace(userTemporaryInformation_.ForgotPasswordVerificationToken) == "" {
			return helper.Translate(lang_, "forgot_password_verification_token_field_empty_validation")
		} else {
			return "ok"
		}
	} else if userTemporaryInformation_.EmailType == "register" {
		if strings.TrimSpace(userTemporaryInformation_.RegisterVerificationToken) == "" {
			return helper.Translate(lang_, "register_verification_token_field_empty_validation")
		} else {
			return "ok"
		}
	} else {
		return "ok"
	}
}
