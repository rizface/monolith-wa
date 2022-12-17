package app

import (
	"database/sql"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/rizface/monolith-mini-whatsapp/db"
	"github.com/rizface/monolith-mini-whatsapp/helper"
)

type App struct {
	Router        *fiber.App
	Postgres      *sql.DB
	Redis         *redis.Client
	Logger        *helper.Logger
	KafkaProducer *kafka.Producer
}

func Init() *App {
	helper.ConfigLoader("./", "config", "env")
	router := NewRouter()
	postgresConnection := db.InitPostgresql()
	redisConnection := db.StartRedisConnection()
	logger := helper.InitNewLogger()
	kafkaProducer := NewProducer()

	return &App{
		Router:        router,
		Postgres:      postgresConnection,
		Logger:        logger,
		Redis:         redisConnection,
		KafkaProducer: kafkaProducer,
	}
}

func (app *App) Start() {
	end := make(chan bool)
	go func() {
		app.Router.Listen(":8000")
	}()
	go StartUserRegisterConsumer(app.Postgres)
	<-end
}
