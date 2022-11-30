package domain

import (
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
)

type UserRequestDomain struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *UserRequestDomain) Validate() []string {
	user := *u
	err := validation.ValidateStruct(
		&user,
		validation.Field(&user.Name, validation.Required),
		validation.Field(&user.Username, validation.Required),
		validation.Field(&user.Password, validation.Required, validation.Length(8, 50)),
	)
	if err == nil {
		return []string{}
	}
	return strings.Split(err.Error(), "; ")
}
