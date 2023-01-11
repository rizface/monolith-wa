package port

import (
	"context"
	"database/sql"

	"github.com/rizface/monolith-mini-whatsapp/db/entity"
	"github.com/rizface/monolith-mini-whatsapp/protocol/domain"
)

type UserRepositoryInterface interface {
	FindByUsername(ctx context.Context, db *sql.DB, username string) (*entity.User, error)
	FindById(ctx context.Context, db *sql.DB, id string) (*entity.User, error)
	Create(ctx context.Context, db *sql.DB, userdomain *domain.UserRequestDomain) error
}
