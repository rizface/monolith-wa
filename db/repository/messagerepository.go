package repository

import (
	"database/sql"

	"github.com/rizface/monolith-mini-whatsapp/protocol/domain"
)

type MessageRepository struct{}

func InitMessageRepository() *MessageRepository {
	return &MessageRepository{}
}

func (m *MessageRepository) Create(db *sql.DB, message *domain.MessageRequestDomain) error {
	sql := "INSERT INTO messages (sender_id, receiver_id, message) VALUES($1, $2, $3)"
	_, err := db.Exec(sql, message.SenderId, message.ReceiverId, message.Message)
	return err
}
