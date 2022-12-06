package helper

import (
	"github.com/dchest/uniuri"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func Marshaller(id, code, message, data interface{}) fiber.Map {
	return fiber.Map{
		"id":      id,
		"code":    code,
		"message": message,
		"data":    data,
	}
}

func ErrMarshaller(code, msg interface{}, logger *zap.Logger) fiber.Map {
	errId := uniuri.New()
	errMap := Marshaller(
		errId,
		code,
		msg,
		nil,
	)

	if logger != nil {
		logger.Error(errId, zap.Any("errData", errMap))
	}
	return errMap
}
