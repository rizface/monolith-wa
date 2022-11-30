package main

import (
	"github.com/rizface/monolith-mini-whatsapp/app"
	"github.com/rizface/monolith-mini-whatsapp/protocol/rest"
)

func main() {
	app := app.Init()

	rest.StartHealthCheckHandler(app)
	rest.StartAuthHandler(app)

	app.Start()
}
