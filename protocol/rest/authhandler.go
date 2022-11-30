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
	redisrepository := repository.InitRedisRepository()
	service := provider.InitAuthService(
		userrepository,
		redisrepository,
		app.Postgres,
		app.Redis,
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
		return c.Status(http.StatusBadRequest).JSON(errMap)
	}

	user, errCreateUser := h.authservice.Register(userdomain)
	if errCreateUser != nil {
		// Handle Error 500 more smooth e.g insert to jaeger

		//===================

		errMap := helper.ErrMarshaller(
			errCreateUser.Code,
			errCreateUser.Message,
			h.app.Logger.Logger,
		)
		return c.Status(errCreateUser.Code).JSON(errMap)
	}

	// RETURN
	return c.Status(200).JSON(helper.Marshaller(
		nil,
		http.StatusOK,
		"success",
		user.ConvertToResponseDomain(),
	))
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	userdomain := new(domain.UserRequestDomain)
	err := json.NewDecoder(bytes.NewReader(c.Body())).Decode(userdomain)
	if err != nil {
		errBuilder := constant.InternalServerError(err.Error())
		errMap := helper.ErrMarshaller(
			errBuilder.Code,
			errBuilder.Message,
			h.app.Logger.Logger,
		)
		return c.Status(errBuilder.Code).JSON(errMap)
	}

	result, errBuilder := h.authservice.Login(userdomain)
	if errBuilder != nil {
		errMap := helper.ErrMarshaller(
			errBuilder.Code,
			errBuilder.Message,
			h.app.Logger.Logger,
		)
		return c.Status(errBuilder.Code).JSON(errMap)
	}
	return c.JSON(result)
}
