package service

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	repo "github.com/salihkemaloglu/gignox-rr-beta-001/repositories"
)

//CreateTokenEndpoint user token creation
func CreateTokenEndpointService(user_ repo.User) (string,error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"username": user_.Username,
		"password": user_.Password,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "",err
	}
	return tokenString,nil
}

//ValidateMiddleware token validation
func ValidateMiddlewareService(authorizationToken_ string) (string, error) {

	if authorizationToken_ != "" {
		token, err := jwt.Parse(authorizationToken_, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			return []byte("secret"), nil
		})
		if err != nil {
			return "", err
		}
		if token.Valid {
			return "ok", nil
		}
		return "Invalid authorization token", nil

	}
	return "An authorization header is required", nil

}