package repository

import (
	"database/sql"

	"github.com/rizface/monolith-mini-whatsapp/db/entity"
)

type UserRegisterEventRepository struct{}

func NewUserRegisterEventRepository() *UserRegisterEventRepository {
	return &UserRegisterEventRepository{}
}

func (u *UserRegisterEventRepository) Create(db *sql.DB, data *entity.User) {
	sql := "INSERT INTO user_register_events (user_id) VALUES($1)"
	db.Exec(sql, data.Id)
}
