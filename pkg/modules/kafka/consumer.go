package kafka

import (
	"core-banking/config"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
)

type ConsumerArgs struct {
	GroupID    *string
	Server     *string
	Offset     *string
	AutoCommit *bool
}

type Consumer struct {
	k *kafka.Reader
}

func NewConsumer(topic string) Consumer {
	cfgKafka, _ := config.NewKafkaConfig(".env")

	brokers := strings.Split(cfgKafka.Host, ",")
	cfg := kafka.ReaderConfig{
		Brokers:         brokers,
		GroupID:         cfgKafka.GroupID,
		Topic:           topic,
		MinBytes:        10e3,            // 10KB
		MaxBytes:        10e6,            // 10MB
		MaxWait:         1 * time.Second, // Maximum amount of time to wait for new data to come when fetching batches of messages from kafka.
		ReadLagInterval: -1,
	}

	reader := kafka.NewReader(cfg)
	return Consumer{
		k: reader,
	}
}
