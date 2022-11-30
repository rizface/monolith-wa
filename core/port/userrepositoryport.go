package port

import (
	"database/sql"

	"github.com/rizface/monolith-mini-whatsapp/db/entity"
	"github.com/rizface/monolith-mini-whatsapp/protocol/domain"
)

type UserRepositoryInterface interface {
	FindByUsername(db *sql.DB, username string) (*entity.User, error)
	Create(db *sql.DB, userdomain *domain.UserRequestDomain) error
}
