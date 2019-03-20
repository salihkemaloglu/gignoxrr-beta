package validation

import (
	"strings"
	db "github.com/salihkemaloglu/gignox-rr-beta-001/mongodb"
)

func UserRegisterFieldValidation(user db.User) string{

	if strings.TrimSpace(user.Username) == "" {
		return "Username can nat be empty"
	} else if strings.TrimSpace(user.Email) == "" {
		return "Email can not be empty"
	} else if strings.TrimSpace(user.Password) == "" {
		return "Password can not be empty"
	} else {
		return "ok"
	}
}

func UserLoginFieldValidation(user db.User) string{

	if strings.TrimSpace(user.Username) == "" {
		return "Username can nat be empty"
	} else if strings.TrimSpace(user.Password) == "" {
		return "Password can not be empty"
	} else {
		return "ok"
	}
}