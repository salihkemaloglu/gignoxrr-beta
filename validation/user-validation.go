package validation

import (
	"strings"
	db "github.com/salihkemaloglu/gignox-rr-beta-001/mongodb"
	gotrans "github.com/salihkemaloglu/gignox-rr-beta-001/services"
)

func UserRegisterFieldValidation(user db.User,lang string) string{
	
	if strings.TrimSpace(user.Username) == "" {
		return gotrans.Translate(lang,"Username_Field_Empty_Validation")
	} else if strings.TrimSpace(user.Email) == "" {
		return gotrans.Translate(lang,"Email_Field_Empty_Validation")
	} else if strings.TrimSpace(user.Password) == "" {
		return gotrans.Translate(lang,"Password_Field_Empty_Validation")
	} else {
		return "ok"
	}
}

func UserLoginFieldValidation(user db.User,lang string) string{

	if strings.TrimSpace(user.Username) == "" {
		return gotrans.Translate(lang,"Username_Field_Empty_Validation")
	} else if strings.TrimSpace(user.Password) == "" {
		return gotrans.Translate(lang,"Password_Field_Empty_Validation")
	} else {
		return "ok"
	}
}