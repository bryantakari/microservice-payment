package kafka

import (
	"context"
	"encoding/json"

	"github.com/bryantakari/microservice-payment/order-service/internal/config"
	"github.com/bryantakari/microservice-payment/order-service/internal/order"
	"github.com/segmentio/kafka-go"
)

type OrderPublisher struct {
	writer *kafka.Writer
}

func NewOrderPublisher(cfg config.KafkaConfig) *OrderPublisher {
	return &OrderPublisher{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(cfg.Brokers...),
			Topic:    cfg.Topics.OrderCreated.Name,
			Balancer: &kafka.LeastBytes{},
		},
	}
}
func (p *OrderPublisher) PublishOrderCreated(
	ctx context.Context,
	event order.OrderCreatedEvent,
) error {

	payload, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return p.writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(event.OrderID),
		Value: payload,
	})
}
