package services

import (
	"context"
	"fmt"

	"github.com/patrickmn/go-cache"
	helper "github.com/salihkemaloglu/gignoxrr-beta-001/helpers"
	inter "github.com/salihkemaloglu/gignoxrr-beta-001/interfaces"
	gigxRR "github.com/salihkemaloglu/gignoxrr-beta-001/proto"
	repo "github.com/salihkemaloglu/gignoxrr-beta-001/repositories"
	val "github.com/salihkemaloglu/gignoxrr-beta-001/validations"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

//LoginService ...
func LoginService(ctx context.Context, req *gigxRR.LoginUserRequest, c *cache.Cache) (*gigxRR.LoginUserResponse, error) {

	userLang := "en"
	if headers, ok := metadata.FromIncomingContext(ctx); ok {
		if headers["languagecode"] != nil {
			userLang = headers["languagecode"][0]
		}
	}
	lang := helper.DetectLanguage(userLang)

	userIPInformationReq, ipErr := helper.GetIPInformation(ctx, false)
	if ipErr != nil {
		return nil, status.Errorf(
			codes.Unavailable,
			fmt.Sprintf(ipErr.Error()),
		)
	}

	userIPInformation := userIPInformationReq.GetIpInformation()

	data := req.GetUser()
	user := repo.User{
		Username: data.GetUsername(),
		Password: data.GetPassword(),
	}
	if res := val.UserLoginFieldValidation(user, lang); res != "ok" {
		return nil, status.Errorf(
			codes.FailedPrecondition,
			fmt.Sprintf(res),
		)
	}
	var loginAttemptCount int
	if x, found := c.Get(userIPInformation.IpAddress); found {
		loginAttemptCount = x.(int)
		if loginAttemptCount >= 20 {
			return nil, status.Errorf(
				codes.Aborted,
				fmt.Sprintf(helper.Translate(lang, "user_login_attemps")),
			)
		}
	}
	user.Password = helper.EncryptePassword(user.Password)
	var op inter.IUserRepository = user
	userResp, err := op.Login()
	if err != nil {
		if x, found := c.Get(userIPInformation.IpAddress); found {
			loginAttemptCount = x.(int)
			if loginAttemptCount < 20 {
				loginAttemptCount = loginAttemptCount + 1
				c.Set(userIPInformation.IpAddress, loginAttemptCount, cache.DefaultExpiration)
			}
		} else {
			c.Set(userIPInformation.IpAddress, 1, cache.DefaultExpiration)
		}
		return nil, status.Errorf(
			codes.Unauthenticated,
			fmt.Sprintf(helper.Translate(lang, "invalid_user_information")),
		)
	}
	c.Delete(userIPInformation.IpAddress)
	tokenRes, tokenErr := helper.CreateTokenEndpointService(user)
	if tokenErr != nil {
		return nil, status.Errorf(
			codes.Unknown,
			fmt.Sprintf(helper.Translate(lang, "token_create_error")+": %v", tokenErr.Error()),
		)
	}

	tokenQC, errQC := helper.GetUserToken(user)

	if errQC != nil {
		return nil, status.Errorf(
			codes.Unknown,
			fmt.Sprintf(helper.Translate(lang, "token_create_error")+": %v", tokenErr.Error()),
		)
	}

	return &gigxRR.LoginUserResponse{
		User: &gigxRR.UserLogin{
			Username:     userResp.Username,
			LanguageCode: userResp.LanguageCode,
			TokenRR:      tokenRes,
			TokenQC:      tokenQC,
		},
	}, nil

}
