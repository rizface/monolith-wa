package app

import (
	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
)

func NewRouter() *fiber.App {
	r := fiber.New()
	r.Use(otelfiber.Middleware("monolith-wa", otelfiber.WithSpanNameFormatter(func(ctx *fiber.Ctx) string {
		return ctx.Path()
	})))
	return r
}
