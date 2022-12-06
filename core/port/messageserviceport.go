package port

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizface/monolith-mini-whatsapp/constant"
	"github.com/rizface/monolith-mini-whatsapp/protocol/domain"
)

type MessageServicePort interface {
	Create(message *domain.MessageRequestDomain) (fiber.Map, *constant.ErrorBuilder)
}
