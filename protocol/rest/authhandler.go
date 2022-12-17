package rest

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rizface/monolith-mini-whatsapp/app"
	"github.com/rizface/monolith-mini-whatsapp/constant"
	"github.com/rizface/monolith-mini-whatsapp/core/port"
	"github.com/rizface/monolith-mini-whatsapp/core/provider"
	"github.com/rizface/monolith-mini-whatsapp/db/repository"
	"github.com/rizface/monolith-mini-whatsapp/helper"
	"github.com/rizface/monolith-mini-whatsapp/protocol/domain"
)

type AuthHandler struct {
	app         *app.App
	authservice port.AuthServicePort
}

func initAuthHandler(app *app.App, authservice port.AuthServicePort) *AuthHandler {
	return &AuthHandler{
		app,
		authservice,
	}
}

func StartAuthHandler(app *app.App) {
	userrepository := repository.InitUserRepository()
	pubsubrepository := repository.NewPubSubRepositoru(app.KafkaProducer)
	redisrepository := repository.InitRedisRepository()
	service := provider.InitAuthService(
		userrepository,
		redisrepository,
		app.Postgres,
		app.Redis,
		pubsubrepository,
	)
	handler := initAuthHandler(app, service)

	app.Router.Post("/v1/register", handler.Register)
	app.Router.Post("/v1/login", handler.Login)
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	userdomain := new(domain.UserRequestDomain)
	err := json.NewDecoder(bytes.NewReader(c.Body())).Decode(userdomain)
	if err != nil {
		errMap := helper.ErrMarshaller(
			http.StatusBadRequest,
			err.Error(),
			h.app.Logger.Logger,
		)
		return c.Status(http.StatusBadRequest).JSON(errMap)
	}

	validationError := userdomain.Validate()
	if len(validationError) > 0 {
		errMap := helper.ErrMarshaller(
			http.StatusBadRequest,
			validationError,
			h.app.Logger.Logger,
		)
		c.Locals("result", errMap)
		return c.Next()
	}

	user, errCreateUser := h.authservice.Register(userdomain)
	if errCreateUser != nil {
		c.Locals("result", errCreateUser)
		return c.Next()
	}

	c.Locals("result", helper.Marshaller(
		nil,
		http.StatusOK,
		"success",
		user.ConvertToResponseDomain(),
	))
	return c.Next()
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	userdomain := new(domain.UserRequestDomain)
	err := json.NewDecoder(bytes.NewReader(c.Body())).Decode(userdomain)
	if err != nil {
		errBuilder := constant.InternalServerError(err.Error())
		c.Locals("result", errBuilder)
		return c.Next()
	}

	result, errBuilder := h.authservice.Login(userdomain)
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
