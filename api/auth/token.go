package auth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/nitinda/microservice-change-log/config"
)

func CreateToken(user_id uint32) (string, error) {
	claims := jwt.MapClaims{}

	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	ss, err := token.SignedString(config.SECRETKEY)

	fmt.Printf("%v %v", ss, err)

	return ss, err
}
