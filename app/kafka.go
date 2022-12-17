package app

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rizface/monolith-mini-whatsapp/db/entity"
	"github.com/rizface/monolith-mini-whatsapp/db/repository"
)

func NewProducer() *kafka.Producer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	})

	if err != nil {
		log.Fatal(err)
	}

	return p
}

func NewConsumer() *kafka.Consumer {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "monolith-wa",
	})

	if err != nil {
		log.Fatal(err)
	}

	return c
}

func StartUserRegisterConsumer(db *sql.DB) {
	userRegisterEventRepository := repository.NewUserRegisterEventRepository()
	topic := "user.register"
	consumer := NewConsumer()

	consumer.SubscribeTopics([]string{topic}, nil)
	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil && err.(kafka.Error).Code() != kafka.ErrTimedOut {
			fmt.Println(err.Error())
		} else {
			userEntity := new(entity.User)
			json.Unmarshal(msg.Value, userEntity)
			userRegisterEventRepository.Create(db, userEntity)
		}
	}
}
