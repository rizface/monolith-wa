package app

import (
	"github.com/gofiber/fiber/v2"
)

func NewRouter() *fiber.App {
	r := fiber.New()
	return r
}
