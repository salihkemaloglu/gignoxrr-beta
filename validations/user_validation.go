package validation

import (
	"strings"
	str "github.com/salihkemaloglu/gignox-rr-beta-001/repositories"
	helper "github.com/salihkemaloglu/gignox-rr-beta-001/services"
)

func UserRegisterFieldValidation(user str.User,lang string) string{
	
	if strings.TrimSpace(user.Username) == "" {
		return helper.Translate(lang,"Username_Field_Empty_Validation")
	} else if strings.TrimSpace(user.Email) == "" {
		return helper.Translate(lang,"Email_Field_Empty_Validation")
	} else if strings.TrimSpace(user.Password) == "" {
		return helper.Translate(lang,"Password_Field_Empty_Validation")
	} else {
		return "ok"
	}
}

func UserLoginFieldValidation(user str.User,lang string) string{

	if strings.TrimSpace(user.Username) == "" {
		return helper.Translate(lang,"Username_Field_Empty_Validation")
	} else if strings.TrimSpace(user.Password) == "" {
		return helper.Translate(lang,"Password_Field_Empty_Validation")
	} else {
		return "ok"
	}
}