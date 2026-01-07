package kafka

import (
	"log"

	"github.com/bryantakari/microservice-payment/order-service/internal/config"
	"github.com/segmentio/kafka-go"
)

func BootstrapKafka(cfg *config.KafkaConfig) {
	conn, err := kafka.Dial("tcp", cfg.Brokers[0])
	if err != nil {
		log.Fatalf("kafka connection failed: %v", err)
	}
	defer conn.Close()

	topic := cfg.Topics.OrderCreated

	err = conn.CreateTopics(kafka.TopicConfig{
		Topic:             topic.Name,
		NumPartitions:     topic.Partitions,
		ReplicationFactor: topic.ReplicationFactor,
	})
	if err != nil {
		log.Printf("topic create skipped or failed: %v", err)
	}
}
