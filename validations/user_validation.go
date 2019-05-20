package validation

import (
	"strings"
	str "github.com/salihkemaloglu/gignox-rr-beta-001/repositories"
	helper "github.com/salihkemaloglu/gignox-rr-beta-001/services"
	zxcvbn "github.com/nbutton23/zxcvbn-go"
)

func UserRegisterFieldValidation(user str.User,lang string) string{
	
	if strings.TrimSpace(user.Username) == "" {
		return helper.Translate(lang,"username_field_empty_validation")
	} else if strings.TrimSpace(user.Email) == "" {
		return helper.Translate(lang,"email_field_empty_validation")
	} else if strings.TrimSpace(user.Password) == "" {
		return helper.Translate(lang,"password_field_empty_validation")
	} else if zxcvbn.PasswordStrength(user.Password,nil ).Score < 2 {
		return helper.Translate(lang,"register_password_strenght_validation_info")
	} else {
		return "ok"
	}
}

func UserLoginFieldValidation(user str.User,lang string) string{

	if strings.TrimSpace(user.Username) == "" {
		return helper.Translate(lang,"username_field_empty_validation")
	} else if strings.TrimSpace(user.Password) == "" {
		return helper.Translate(lang,"password_field_empty_validation")
	} else {
		return "ok"
	}
}