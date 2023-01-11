package middleware

import (
	"errors"
	"net/http"

	"github.com/dchest/uniuri"
	"github.com/gofiber/fiber/v2"
	"github.com/rizface/monolith-mini-whatsapp/app"
	"github.com/rizface/monolith-mini-whatsapp/constant"
	"go.opentelemetry.io/otel/codes"
	"go.uber.org/zap"
)

type LoggerHandler struct {
	logger *zap.Logger
}

func initLoggerHandler(app *app.App) *LoggerHandler {
	return &LoggerHandler{
		logger: app.Logger.Logger,
	}
}

func UseLoggerHandler(app *app.App) {
	handler := initLoggerHandler(app)
	app.Router.Use(handler.Log)
}

func (h *LoggerHandler) Log(c *fiber.Ctx) error {
	_, span := tracer.Start(c.UserContext(), "log handler")
	defer span.End()

	result := c.Locals("result")
	errBuilder, ok := result.(*constant.ErrorBuilder)
	if ok {
		id := uniuri.New()
		if errBuilder.Code == http.StatusInternalServerError {
			h.logger.Error(
				id,
				zap.Any("errData", errBuilder),
			)
		}
		span.RecordError(errors.New(errBuilder.Message))
		span.SetStatus(codes.Code(errBuilder.Code), "")
		return c.Status(errBuilder.Code).JSON(fiber.Map{
			"code":    errBuilder.Code,
			"message": errBuilder.Message,
			"data":    nil,
			"id":      id,
		})
	}

	errMap, ok := result.(fiber.Map)
	if ok {
		code := errMap["code"].(int)
		if code >= http.StatusInternalServerError {
			h.logger.Error(
				errMap["id"].(string),
				zap.Any("errData", errMap),
			)
			errMap["message"] = "INTERNAL_SERVER_ERROR"
		}
		span.RecordError(errors.New(errMap["message"].(string)))
		span.SetStatus(codes.Code(errMap["code"].(int)), "")
		return c.Status(code).JSON(errMap)
	}
	return c.JSON(result)
}
