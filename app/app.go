package app

import (
	"database/sql"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/rizface/monolith-mini-whatsapp/db"
	"github.com/rizface/monolith-mini-whatsapp/helper"
)

type App struct {
	Router   *fiber.App
	Postgres *sql.DB
	Redis    *redis.Client
	Logger   *helper.Logger
}

func Init() *App {
	helper.ConfigLoader("./", "config", "env")
	router := NewRouter()
	postgresConnection := db.InitPostgresql()
	redisConnection := db.StartRedisConnection()
	logger := helper.InitNewLogger()

	return &App{
		Router:   router,
		Postgres: postgresConnection,
		Logger:   logger,
		Redis:    redisConnection,
	}
}

func (app *App) Start() {
	err := app.Router.Listen(":8000")
	if err != nil {
		log.Fatal(err)
	}
}
