package rest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/rizface/monolith-mini-whatsapp/app"
	"github.com/rizface/monolith-mini-whatsapp/constant"
	"github.com/rizface/monolith-mini-whatsapp/core/port"
	"github.com/rizface/monolith-mini-whatsapp/core/provider"
	"github.com/rizface/monolith-mini-whatsapp/db/repository"
	"github.com/rizface/monolith-mini-whatsapp/helper"
	"github.com/rizface/monolith-mini-whatsapp/protocol/domain"
	"github.com/rizface/monolith-mini-whatsapp/protocol/rest/middleware"
	"github.com/rizface/monolith-mini-whatsapp/protocol/rest/validator"
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
		router.Get("/sender/:senderId/receiver/:receiverId", handler.GetMessages)
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

func (h *MessageHandler) GetMessages(c *fiber.Ctx) error {
	err := validator.ValidateUUID(c.Params("senderId"))
	if err != nil {
		c.Locals("result", helper.ErrMarshaller(
			http.StatusBadRequest,
			[]string{err.Error()},
			nil,
		))
		return c.Next()
	}

	err = validator.ValidateUUID(c.Params("receiverId"))
	if err != nil {
		c.Locals("result", helper.ErrMarshaller(
			http.StatusBadRequest,
			[]string{err.Error()},
			nil,
		))
		return c.Next()
	}

	senderId, receiverId := c.Params("senderId"), c.Params("receiverId")
	userData := new(helper.Claim)
	err = json.NewDecoder(strings.NewReader(c.Get("USER-DATA"))).Decode(userData)
	if err != nil {
		c.Locals("result", constant.InternalServerError(err.Error()))
		return c.Next()
	}
	messages, errBuilder := h.Service.GetMessages(senderId, receiverId, userData)
	if errBuilder != nil {
		c.Locals("result", errBuilder)
	} else {
		c.Locals("result", helper.Marshaller(
			nil,
			http.StatusOK,
			"SUCCESS",
			messages,
		))
	}
	return c.Next()
}
