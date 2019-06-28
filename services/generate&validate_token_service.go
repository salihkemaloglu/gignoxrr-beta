package service

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	repo "github.com/salihkemaloglu/gignoxrr-beta-001/repositories"
)

//CreateTokenEndpointService user token creation
func CreateTokenEndpointService(user repo.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"username": user.Username,
		"password": user.Password,
		"exp":      time.Now().Add(time.Hour * 571).Unix(),
	})
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

//ValidateMiddlewareService token validation
func ValidateMiddlewareService(authorizationToken string) (string, error) {

	if authorizationToken != "" {
		token, err := jwt.Parse(authorizationToken, func(token *jwt.Token) (interface{}, error) {
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
