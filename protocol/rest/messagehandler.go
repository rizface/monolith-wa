package rest

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rizface/monolith-mini-whatsapp/app"
	"github.com/rizface/monolith-mini-whatsapp/core/port"
	"github.com/rizface/monolith-mini-whatsapp/core/provider"
	"github.com/rizface/monolith-mini-whatsapp/db/repository"
	"github.com/rizface/monolith-mini-whatsapp/helper"
	"github.com/rizface/monolith-mini-whatsapp/protocol/domain"
	"github.com/rizface/monolith-mini-whatsapp/protocol/rest/middleware"
	"go.uber.org/zap"
)

type MessageHandler struct {
	Service port.MessageServicePort
	Logger  *zap.Logger
}

func initMessageHandler(app *app.App, service port.MessageServicePort) *MessageHandler {
	return &MessageHandler{
		Service: service,
		Logger:  app.Logger.Logger,
	}
}

func StartMessageHandler(app *app.App) {
	userRepo := repository.InitUserRepository()
	messageRepo := repository.InitMessageRepository()
	service := provider.InitMessageService(app.Postgres, userRepo, messageRepo)
	handler := initMessageHandler(app, service)
	authHandler := middleware.InitHandler(app)

	app.Router.Route("/v1/messages", func(router fiber.Router) {
		router.Use(authHandler.Auth)
		router.Post("", handler.Create)
	})
}

func (m *MessageHandler) Create(c *fiber.Ctx) error {
	message := new(domain.MessageRequestDomain)
	err := json.NewDecoder(bytes.NewReader(c.Request().Body())).Decode(message)
	if err != nil {
		errMap := helper.ErrMarshaller(
			http.StatusInternalServerError,
			err.Error(),
			m.Logger,
		)
		c.Locals("result", errMap)
		return c.Next()
	}

	errors := message.Validate()
	if len(errors) > 0 {
		errMap := helper.ErrMarshaller(
			http.StatusBadRequest,
			errors,
			nil,
		)
		c.Locals("result", errMap)
		return c.Next()
	}

	result, errBuilder := m.Service.Create(message)
	if errBuilder != nil {
		c.Locals("result", errBuilder)
		return c.Next()
	}

	c.Locals("result", helper.Marshaller(
		nil,
		http.StatusOK,
		"SUCCESS",
		result,
	))

	return c.Next()
}
