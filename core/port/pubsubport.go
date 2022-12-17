package port

type PubSubPort interface {
	Produce(topic string, payload []byte)
}
