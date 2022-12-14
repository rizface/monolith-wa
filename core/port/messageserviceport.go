package port

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/rizface/monolith-mini-whatsapp/constant"
	"github.com/rizface/monolith-mini-whatsapp/db/entity"
	"github.com/rizface/monolith-mini-whatsapp/helper"
	"github.com/rizface/monolith-mini-whatsapp/protocol/domain"
)

type MessageServicePort interface {
	Create(ctx context.Context, message *domain.MessageRequestDomain, userData *helper.Claim) (fiber.Map, *constant.ErrorBuilder)
	GetMessages(ctx context.Context, senderId string, receiverId string, userData *helper.Claim) (*[]entity.Message, *constant.ErrorBuilder)
}
