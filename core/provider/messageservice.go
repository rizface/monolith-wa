package provider

import (
	"database/sql"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/rizface/monolith-mini-whatsapp/constant"
	"github.com/rizface/monolith-mini-whatsapp/core/port"
	"github.com/rizface/monolith-mini-whatsapp/protocol/domain"
)

type MessageService struct {
	db          *sql.DB
	userRepo    port.UserRepositoryInterface
	messageRepo port.MessageRepositoryPort
}

func InitMessageService(
	db *sql.DB,
	userRepo port.UserRepositoryInterface,
	messageRepo port.MessageRepositoryPort,
) *MessageService {
	return &MessageService{
		db:          db,
		userRepo:    userRepo,
		messageRepo: messageRepo,
	}
}

func (m *MessageService) Create(message *domain.MessageRequestDomain) (fiber.Map, *constant.ErrorBuilder) {
	_, err := m.userRepo.FindById(m.db, message.SenderId)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, constant.InternalServerError(err.Error())
		}
		return nil, constant.USER_NOT_FOUND
	}

	_, err = m.userRepo.FindById(m.db, message.ReceiverId)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, constant.InternalServerError(err.Error())
		}
		return nil, constant.USER_NOT_FOUND
	}

	err = m.messageRepo.Create(m.db, message)
	if err != nil {
		return nil, constant.InternalServerError(err.Error())
	}

	return nil, nil
}
