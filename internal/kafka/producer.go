package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// Producer representa un productor Kafka
type Producer struct {
	producer *kafka.Producer
}

// NewProducer crea una nueva instancia del productor Kafka
func NewProducer(bootstrapServers string) (*Producer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": bootstrapServers})
	if err != nil {
		return nil, err
	}

	return &Producer{producer: p}, nil
}

// Produce envía un mensaje al topic especificado
func (p *Producer) Produce(topic string, message []byte) error {
	deliveryChan := make(chan kafka.Event)
	err := p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: -1},
		Value:          message,
	}, deliveryChan)

	if err != nil {
		return err
	}

	// Espera la entrega del mensaje y maneja el resultado
	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		return m.TopicPartition.Error
	}

	return nil
}

// Close cierra el productor Kafka
func (p *Producer) Close() {
	p.producer.Close()
}
