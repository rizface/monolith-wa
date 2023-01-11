package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/rizface/monolith-mini-whatsapp/db/entity"
	"github.com/rizface/monolith-mini-whatsapp/protocol/domain"
)

type MessageRepository struct{}

func InitMessageRepository() *MessageRepository {
	return &MessageRepository{}
}

func (m *MessageRepository) Create(ctx context.Context, db *sql.DB, message *domain.MessageRequestDomain) error {
	ctx, span := tracer.Start(ctx, "messagerepository.Create")
	defer span.End()
	sql := "INSERT INTO messages (sender_id, receiver_id, message) VALUES($1, $2, $3)"
	_, err := db.Exec(sql, message.SenderId, message.ReceiverId, message.Message)
	return err
}

func (m *MessageRepository) GetMessages(ctx context.Context, db *sql.DB, senderId string, receiverId string) (*[]entity.Message, error) {
	ctx, span := tracer.Start(ctx, "messagerepository.GetMessages")
	defer span.End()

	sql := `
		SELECT id, sender_id, receiver_id, message, created_at, updated_at FROM messages
		WHERE sender_id = $1 AND receiver_id = $2 OR sender_id = $2 AND receiver_id = $1
		ORDER BY created_at ASC
	`
	rows, err := db.Query(sql, senderId, receiverId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	messages := []entity.Message{}
	for rows.Next() {
		message := entity.Message{}
		err := rows.Scan(
			&message.Id,
			&message.SenderId,
			&message.ReceiverId,
			&message.Message,
			&message.CreatedAt,
			&message.UpdatedAt,
		)

		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		messages = append(messages, message)
	}
	return &messages, nil
}
