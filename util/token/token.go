package token

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/globalsign/mgo/bson"
	"log-collector/config"
)

var tokenSecret = []byte(config.UserTokenSecret)

func Generate(userID bson.ObjectId) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": userID.Hex(),
	})

	return token.SignedString(tokenSecret)
}

func Validate(tokenString string) (jwt.MapClaims, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return tokenSecret, nil
	})
	if err != nil || !token.Valid {
		return nil, false
	}

	claims := token.Claims.(jwt.MapClaims)
	return claims, true
}
