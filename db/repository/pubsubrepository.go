package repository

import "github.com/confluentinc/confluent-kafka-go/kafka"

type PubSubRepository struct {
	producer *kafka.Producer
}

func NewPubSubRepositoru(producer *kafka.Producer) *PubSubRepository {
	return &PubSubRepository{
		producer: producer,
	}
}

func (p *PubSubRepository) Produce(topic string, payload []byte) {
	p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          payload,
	}, nil)
}
