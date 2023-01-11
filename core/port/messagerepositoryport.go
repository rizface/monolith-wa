package port

import (
	"context"
	"database/sql"

	"github.com/rizface/monolith-mini-whatsapp/db/entity"
	"github.com/rizface/monolith-mini-whatsapp/protocol/domain"
)

type MessageRepositoryPort interface {
	Create(ctx context.Context, db *sql.DB, message *domain.MessageRequestDomain) error
	GetMessages(ctx context.Context, db *sql.DB, senderId string, receiverId string) (*[]entity.Message, error)
}
