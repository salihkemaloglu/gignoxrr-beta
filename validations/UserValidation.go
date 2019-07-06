package validations

import (
	"strings"

	zxcvbn "github.com/nbutton23/zxcvbn-go"
	helper "github.com/salihkemaloglu/gignoxrr-beta-001/helpers"
	repo "github.com/salihkemaloglu/gignoxrr-beta-001/repositories"
)

//UserRegisterFieldValidation ...
func UserRegisterFieldValidation(user repo.User, lang string) string {

	if strings.TrimSpace(user.Username) == "" {
		return helper.Translate(lang, "username_field_empty_validation")
	} else if strings.TrimSpace(user.Email) == "" {
		return helper.Translate(lang, "email_field_empty_validation")
	} else if strings.TrimSpace(user.Password) == "" {
		return helper.Translate(lang, "password_field_empty_validation")
	} else if zxcvbn.PasswordStrength(user.Password, nil).Score < 2 {
		return helper.Translate(lang, "register_password_strenght_validation_information")
	} else {
		return "ok"
	}
}

//UserLoginFieldValidation ...
func UserLoginFieldValidation(user repo.User, lang string) string {

	if strings.TrimSpace(user.Username) == "" {
		return helper.Translate(lang, "username_field_empty_validation")
	} else if strings.TrimSpace(user.Password) == "" {
		return helper.Translate(lang, "password_field_empty_validation")
	} else {
		return "ok"
	}
}

//GetUserFieldValidation ...
func GetUserFieldValidation(user repo.User, lang string) string {

	if strings.TrimSpace(user.Username) == "" {
		return helper.Translate(lang, "username_field_empty_validation")
	}
	return "ok"

}
