package provider

import (
	"context"
	"database/sql"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/rizface/monolith-mini-whatsapp/constant"
	"github.com/rizface/monolith-mini-whatsapp/core/port"
	"github.com/rizface/monolith-mini-whatsapp/db/entity"
	"github.com/rizface/monolith-mini-whatsapp/helper"
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

func (m *MessageService) Create(ctx context.Context, message *domain.MessageRequestDomain, userData *helper.Claim) (fiber.Map, *constant.ErrorBuilder) {
	ctx, span := tracer.Start(ctx, "messageservice.Create")
	defer span.End()

	_, err := m.userRepo.FindById(ctx, m.db, message.SenderId)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, constant.InternalServerError(err.Error())
		}
		return nil, constant.USER_NOT_FOUND
	}

	_, err = m.userRepo.FindById(ctx, m.db, message.ReceiverId)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, constant.InternalServerError(err.Error())
		}
		return nil, constant.USER_NOT_FOUND
	}

	if message.SenderId != userData.Id {
		return nil, constant.UNAUTHORIZED
	}

	err = m.messageRepo.Create(ctx, m.db, message)
	if err != nil {
		return nil, constant.InternalServerError(err.Error())
	}

	return nil, nil
}

func (m *MessageService) GetMessages(ctx context.Context, senderId string, receiverId string, userData *helper.Claim) (*[]entity.Message, *constant.ErrorBuilder) {
	ctx, span := tracer.Start(ctx, "messageservice.GetMessages")
	defer span.End()

	sender, err := m.userRepo.FindById(ctx, m.db, senderId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, constant.SENDER_NOT_FOUND
		}
		return nil, constant.InternalServerError(err.Error())
	}

	receiver, err := m.userRepo.FindById(ctx, m.db, receiverId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, constant.RECEIVER_NOT_FOUND
		}
		return nil, constant.InternalServerError(err.Error())
	}
	if sender.Id != userData.Id {
		return nil, constant.UNAUTHORIZED
	}

	messages, err := m.messageRepo.GetMessages(ctx, m.db, sender.Id, receiver.Id)
	if err != nil {
		return nil, constant.InternalServerError(err.Error())
	}
	return messages, nil
}
