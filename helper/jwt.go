package helper

import (
	"time"

	"github.com/dchest/uniuri"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rizface/monolith-mini-whatsapp/db/entity"
	"github.com/spf13/viper"
)

type Claim struct {
	Name     string
	Username string
	jwt.RegisteredClaims
}

func GenerateJWT(user *entity.User) (interface{}, error) {
	claim := Claim{
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
