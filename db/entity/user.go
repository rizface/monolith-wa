package entity

import (
	"time"

	"github.com/rizface/monolith-mini-whatsapp/protocol/domain"
)

type User struct {
	Id        string
	Name      string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) ConvertToResponseDomain() *domain.UserResponseDomain {
	return &domain.UserResponseDomain{
		Id:        u.Id,
		Username:  u.Username,
		Name:      u.Name,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
