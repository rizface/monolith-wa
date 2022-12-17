package port

import (
	"database/sql"

	"github.com/rizface/monolith-mini-whatsapp/db/entity"
)

type UserRegisterEventPort interface {
	Create(db *sql.DB, data *entity.User)
}
