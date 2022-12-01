package middleware

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/rizface/monolith-mini-whatsapp/app"
	"github.com/rizface/monolith-mini-whatsapp/constant"
	"github.com/rizface/monolith-mini-whatsapp/helper"
)

type AuthHandler struct {
	app *app.App
}

func InitHandler(app *app.App) *AuthHandler {
	return &AuthHandler{
		app: app,
	}
}

func (a *AuthHandler) Auth(c *fiber.Ctx) error {
	headerValue := c.Get("Authorization")
	splittedHeaderValue := strings.Split(headerValue, "Bearer")
	if len(splittedHeaderValue) != 2 {
		errBuilder := constant.INVALID_TOKEN
		errMap := helper.ErrMarshaller(
			errBuilder.Code,
			errBuilder.Message,
			a.app.Logger.Logger,
		)
		return c.Status(errBuilder.Code).JSON(errMap)
	}

	token := strings.Trim(splittedHeaderValue[1], " ")
	result := a.app.Redis.Get(c.Context(), token)
	if result.Err() != nil && errors.Is(result.Err(), redis.Nil) {
		errBuilder := constant.INVALID_TOKEN
		errMap := helper.ErrMarshaller(
			errBuilder.Code,
			errBuilder.Message,
			a.app.Logger.Logger,
		)
		return c.Status(errBuilder.Code).JSON(errMap)
	}

	claim, errBuilder := helper.DecodeJWT(token)
	if errBuilder != nil {
		errMap := helper.ErrMarshaller(
			errBuilder.Code,
			errBuilder.Message,
			a.app.Logger.Logger,
		)
		return c.Status(errBuilder.Code).JSON(errMap)
	}

	bytesClaim, _ := json.Marshal(claim)
	c.Request().Header.Add("USER-DATA", string(bytesClaim))
	return c.Next()
}
