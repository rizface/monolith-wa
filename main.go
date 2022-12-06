package main

import (
	"github.com/rizface/monolith-mini-whatsapp/app"
	"github.com/rizface/monolith-mini-whatsapp/protocol/rest"
	"github.com/rizface/monolith-mini-whatsapp/protocol/rest/middleware"
)

func main() {
	app := app.Init()

	rest.StartHealthCheckHandler(app)
	rest.StartAuthHandler(app)
	rest.StartMessageHandler(app)
	middleware.UseLoggerHandler(app)

	app.Start()
}
