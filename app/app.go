package app

import (
	"context"
	"database/sql"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/rizface/monolith-mini-whatsapp/db"
	"github.com/rizface/monolith-mini-whatsapp/helper"
	"go.opentelemetry.io/otel/sdk/trace"
)

type App struct {
	Router         *fiber.App
	Postgres       *sql.DB
	Redis          *redis.Client
	Logger         *helper.Logger
	KafkaProducer  *kafka.Producer
	TracerProvider *trace.TracerProvider
}

func Init() *App {
	helper.ConfigLoader("./", "config", "env")
	router := NewRouter()
	postgresConnection := db.InitPostgresql()
	redisConnection := db.StartRedisConnection()
	logger := helper.InitNewLogger()
	kafkaProducer := NewProducer()
	tracerProvider := NewOtelProvider()

	return &App{
		Router:         router,
		Postgres:       postgresConnection,
		Logger:         logger,
		Redis:          redisConnection,
		KafkaProducer:  kafkaProducer,
		TracerProvider: tracerProvider,
	}
}

func (app *App) Start() {
	end := make(chan bool)
	ctx := context.Background()
	defer app.TracerProvider.Shutdown(ctx)

	go func() {
		app.Router.Listen(":8000")
	}()
	go StartUserRegisterConsumer(app.Postgres)
	<-end
}
