package jwt

import (
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func Create(userName string) (string, error) {
	signBytes, err := ioutil.ReadFile("game/app/pkg/jwt/test.rsa")
	if err != nil {
		return "", err
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = userName
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	tokenString, err := token.SignedString(signBytes)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
