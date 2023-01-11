package port

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/rizface/monolith-mini-whatsapp/constant"
	"github.com/rizface/monolith-mini-whatsapp/db/entity"
	"github.com/rizface/monolith-mini-whatsapp/protocol/domain"
)

type AuthServicePort interface {
	Register(ctx context.Context, userdomain *domain.UserRequestDomain) (*entity.User, *constant.ErrorBuilder)
	Login(ctx context.Context, userdomain *domain.UserRequestDomain) (fiber.Map, *constant.ErrorBuilder)
}
