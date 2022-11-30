package rest

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rizface/monolith-mini-whatsapp/app"
	"github.com/rizface/monolith-mini-whatsapp/helper"
)

type HealtchCheckHandler struct {
	App *app.App
}

func initHealtchCheckHandler(app *app.App) *HealtchCheckHandler {
	return &HealtchCheckHandler{
		App: app,
	}
}

func StartHealthCheckHandler(app *app.App) {
	handler := initHealtchCheckHandler(app)
	app.Router.Get("/healthcheck", handler.HealthCheck)
}

func test(ctx context.Context) {
	time.Sleep(6 * time.Second)
	return
}

func (h *HealtchCheckHandler) HealthCheck(c *fiber.Ctx) error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	errCh := make(chan error)
	defer cancel()

	go func() {
		errCh <- h.App.Postgres.PingContext(ctx)
	}()

	select {
	case <-ctx.Done():
		return c.Status(http.StatusServiceUnavailable).JSON(helper.Marshaller(
			nil,
			http.StatusServiceUnavailable,
			"App Not Ready, Failed Establish Connection To Database",
			nil,
		))
	case err := <-errCh:
		if err != nil {
			return c.Status(http.StatusServiceUnavailable).JSON(helper.Marshaller(
				nil,
				http.StatusServiceUnavailable,
				err.Error(),
				nil,
			))
		}
		return c.Status(http.StatusOK).JSON(helper.Marshaller(
			nil,
			http.StatusOK,
			"App Ready",
			nil,
		))
	}
}
