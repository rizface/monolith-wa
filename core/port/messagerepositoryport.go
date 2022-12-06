package port

import (
	"database/sql"

	"github.com/rizface/monolith-mini-whatsapp/db/entity"
	"github.com/rizface/monolith-mini-whatsapp/protocol/domain"
)

type MessageRepositoryPort interface {
	Create(db *sql.DB, message *domain.MessageRequestDomain) error
	GetMessages(db *sql.DB, senderId string, receiverId string) (*[]entity.Message, error)
}
