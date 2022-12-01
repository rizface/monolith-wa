package helper

import (
	"errors"
	"time"

	"github.com/dchest/uniuri"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rizface/monolith-mini-whatsapp/constant"
	"github.com/rizface/monolith-mini-whatsapp/db/entity"
	"github.com/spf13/viper"
)

type Claim struct {
	Id       string
	Name     string
	Username string
	jwt.RegisteredClaims
}

func GenerateJWT(user *entity.User) (interface{}, error) {
	claim := Claim{
		Id:       user.Id,
		Name:     user.Name,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(48 * time.Hour)),
			Issuer:    "rizface",
			ID:        uniuri.New(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString([]byte(viper.GetString("JWT_SECRET")))
	if err != nil {
		return nil, err
	}
	return tokenString, nil
}

func DecodeJWT(tokenString string) (*Claim, *constant.ErrorBuilder) {
	token, err := jwt.ParseWithClaims(tokenString, &Claim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("JWT_SECRET")), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, constant.TOKEN_EXPIRED
		}
		return nil, constant.INVALID_TOKEN
	}

	if !token.Valid {
		return nil, constant.INVALID_TOKEN
	}

	claim, ok := token.Claims.(*Claim)
	if !ok {
		return nil, constant.INVALID_TOKEN
	}

	return claim, nil
}
