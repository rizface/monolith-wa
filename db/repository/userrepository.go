package repository

import (
	"database/sql"

	"github.com/rizface/monolith-mini-whatsapp/core/port"
	"github.com/rizface/monolith-mini-whatsapp/db/entity"
	"github.com/rizface/monolith-mini-whatsapp/protocol/domain"
)

type UserRepository struct{}

func InitUserRepository() port.UserRepositoryInterface {
	return &UserRepository{}
}

func (u *UserRepository) FindByUsername(db *sql.DB, username string) (*entity.User, error) {
	user := entity.User{}
	row := db.QueryRow("SELECT id, name, username, password, created_at, updated_at FROM users WHERE username = $1", username)
	err := row.Scan(&user.Id, &user.Name, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepository) FindById(db *sql.DB, id string) (*entity.User, error) {
	user := entity.User{}
	row := db.QueryRow("SELECT id, name, username, password, created_at, updated_at FROM users WHERE id = $1", id)
	err := row.Scan(&user.Id, &user.Name, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepository) Create(db *sql.DB, userdomain *domain.UserRequestDomain) error {
	_, err := db.Exec(
		"INSERT INTO users (name,username,password) VALUES($1,$2,$3)",
		userdomain.Name,
		userdomain.Username,
		userdomain.Password,
	)
	return err
}
