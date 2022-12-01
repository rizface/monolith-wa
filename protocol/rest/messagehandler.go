package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizface/monolith-mini-whatsapp/app"
	"github.com/rizface/monolith-mini-whatsapp/protocol/rest/middleware"
)

type MessageHandler struct{}

func initMessageHandler() *MessageHandler {
	return &MessageHandler{}
}

func StartMessageHandler(app *app.App) {
	handler := initMessageHandler()
	authHandler := middleware.InitHandler(app)

	app.Router.Route("/v1/messages", func(router fiber.Router) {
		router.Use(authHandler.Auth)
		router.Post("", handler.Create)
	})
}

func (m *MessageHandler) Create(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "SUCCESS",
	})
}
