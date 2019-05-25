package validation

import (
	"strings"
	repo "github.com/salihkemaloglu/gignox-rr-beta-001/repositories"
	helper "github.com/salihkemaloglu/gignox-rr-beta-001/services"
	zxcvbn "github.com/nbutton23/zxcvbn-go"
)

func UserRegisterFieldValidation(user_ repo.User,lang_ string) string{
	
	if strings.TrimSpace(user_.Username) == "" {
		return helper.Translate(lang_,"username_field_empty_validation")
	} else if strings.TrimSpace(user_.Email) == "" {
		return helper.Translate(lang_,"email_field_empty_validation")
	} else if strings.TrimSpace(user_.Password) == "" {
		return helper.Translate(lang_,"password_field_empty_validation")
	} else if zxcvbn.PasswordStrength(user_.Password,nil ).Score < 2 {
		return helper.Translate(lang_,"register_password_strenght_validation_information")
	} else {
		return "ok"
	}
}

func UserLoginFieldValidation(user_ repo.User,lang_ string) string{

	if strings.TrimSpace(user_.Username) == "" {
		return helper.Translate(lang_,"username_field_empty_validation")
	} else if strings.TrimSpace(user_.Password) == "" {
		return helper.Translate(lang_,"password_field_empty_validation")
	} else {
		return "ok"
	}
}
func GetUserFieldValidation(user_ repo.User,lang_ string) string{

	if strings.TrimSpace(user_.Username) == "" {
		return helper.Translate(lang_,"username_field_empty_validation")
	} else {
		return "ok"
	}
}