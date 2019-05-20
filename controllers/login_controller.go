package controller

import (
	"fmt"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/metadata"
	"github.com/patrickmn/go-cache"
	"github.com/salihkemaloglu/gignox-rr-beta-001/proto"
	repo "github.com/salihkemaloglu/gignox-rr-beta-001/repositories"
	val "github.com/salihkemaloglu/gignox-rr-beta-001/validations"
	inter "github.com/salihkemaloglu/gignox-rr-beta-001/interfaces"
	helper "github.com/salihkemaloglu/gignox-rr-beta-001/services"
)
func LoginController(ctx_ context.Context, req_ *gigxRR.LoginUserRequest,c *cache.Cache) (*gigxRR.LoginUserResponse, error) {

	userLang :="en"
	if headers, ok := metadata.FromIncomingContext(ctx_); ok {
		userLang = headers["language"][0]
	}
	lang := helper.DetectLanguage(userLang)

	userIpInformationReq,ipErr:=helper.GetIpInformation(ctx_,false)
	if ipErr !=nil {
		return nil,status.Errorf(
			codes.Unavailable,
			fmt.Sprintf(ipErr.Error()),
		)
	}
	
	userIpInformation:=userIpInformationReq.GetIpInformation()

	data := req_.GetUser();
	user := repo.User {
		Username:data.GetUsername(),
		Password:data.GetPassword(),
	}
	if res := val.UserLoginFieldValidation(user,lang); res != "ok" {
		return nil,status.Errorf(
			codes.FailedPrecondition,
			fmt.Sprintf(res),
		)
	}
	var loginAttemptCount int
	if x, found := c.Get(userIpInformation.IpAddress); found {
		loginAttemptCount = x.(int)
		if loginAttemptCount >= 20 {
			return nil,status.Errorf(
				codes.Aborted,
				fmt.Sprintf(helper.Translate(lang,"user_login_attemps")),
			)
		}
	}
	var op inter.IUserRepository=user
	userResp,err:= op.Login();
	if err != nil {
	    if x, found := c.Get(user.Username); found {
			loginAttemptCount = x.(int)
			if loginAttemptCount < 20 {
				loginAttemptCount=loginAttemptCount+1
				c.Set(user.Username, loginAttemptCount, cache.DefaultExpiration)
			}
	    } else {
			c.Set(user.Username, 1, cache.DefaultExpiration)
		}
		return nil,status.Errorf(
			codes.Unauthenticated,
			fmt.Sprintf(helper.Translate(lang,"invalid_user_information")),
		)
	}
	tokenRes,tokenErr:=helper.CreateTokenEndpointService(user)
	if tokenErr != nil{
		return nil,status.Errorf(
			codes.Unknown,
			fmt.Sprintf(helper.Translate(lang,"token_create_error") +": %v",tokenErr.Error()),
		)
	}

	return &gigxRR.LoginUserResponse{
		User:&gigxRR.UserLogin{
			Username:	userResp.Username,
			LanguageCode:userResp.LanguageCode,
			Token:		tokenRes,
		},
	}, nil

}

