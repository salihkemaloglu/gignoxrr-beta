package controller

import (
	"fmt"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/metadata"
	"github.com/patrickmn/go-cache"
	"github.com/salihkemaloglu/gignox-rr-beta-001/proto"
	db "github.com/salihkemaloglu/gignox-rr-beta-001/mongodb"
	val "github.com/salihkemaloglu/gignox-rr-beta-001/validation"
	repo "github.com/salihkemaloglu/gignox-rr-beta-001/repository"
	helper "github.com/salihkemaloglu/gignox-rr-beta-001/services"
)
func LoginController(ctx context.Context, req *gigxRR.LoginUserRequest,c *cache.Cache) (*gigxRR.LoginUserResponse, error) {

	userLang :="en"
	if headers, ok := metadata.FromIncomingContext(ctx); ok {
		fmt.Println(headers["user-agent"][0])
		userLang = headers["language"][0]
	}
	lang := helper.DetectLanguage(userLang)

	data := req.GetUser();
	user := db.User {
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
	if x, found := c.Get(user.Username); found {
		loginAttemptCount = x.(int)
		if loginAttemptCount >= 20{
			return nil,status.Errorf(
				codes.Aborted,
				fmt.Sprintf(helper.Translate(lang,"User_Login_attemps")),
			)
		}
	}
	var op repo.UserRepository=user
	if  err:= op.Login(); err != nil {
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
			fmt.Sprintf(helper.Translate(lang,"Invalid_User_Information")),
		)
	}
	tokenRes,tokenErr:=helper.CreateTokenEndpointService(user)
	if tokenErr != nil{
		return nil,status.Errorf(
			codes.Unknown,
			fmt.Sprintf(helper.Translate(lang,"Token_Create_Error") +": %v",tokenErr.Error()),
		)
	}

	return &gigxRR.LoginUserResponse{
		User:&gigxRR.UserLogin{
			Username:	user.Username,
			Token:		tokenRes,
		},
	}, nil

}

