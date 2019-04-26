package validation

import (
	"strings"
	db "github.com/salihkemaloglu/gignox-rr-beta-001/mongodb"
	helper "github.com/salihkemaloglu/gignox-rr-beta-001/services"
)

func UserRegisterFieldValidation(user db.User,lang string) string{
	
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

func UserLoginFieldValidation(user db.User,lang string) string{

	if strings.TrimSpace(user.Username) == "" {
		return helper.Translate(lang,"Username_Field_Empty_Validation")
	} else if strings.TrimSpace(user.Password) == "" {
		return helper.Translate(lang,"Password_Field_Empty_Validation")
	} else {
		return "ok"
	}
}