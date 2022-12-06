package port

import (
	"database/sql"

	"github.com/rizface/monolith-mini-whatsapp/protocol/domain"
)

type MessageRepositoryPort interface {
	Create(db *sql.DB, message *domain.MessageRequestDomain) error
}
