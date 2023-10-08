package kafka

import (
	"context"
	"core-banking/config"
	"errors"
	"fmt"
	"math"
	"sync"
	"time"

	"github.com/gofrs/uuid"
	"github.com/segmentio/kafka-go"
)

var (
	kafkaWriters = make(map[string]*kafka.Writer)
	syncMutex    sync.Mutex
)

type Producer interface {
	PublishMessage(topic string, messages []string)
	SendMessage(topic string, messages []string)
}

// Producer struct for producer needed
type producerImpl struct {
	k *kafka.Writer
}

// NewProducer Create  for segmentio kafka
func NewProducer() Producer {
	cfgKafka, _ := config.NewKafkaConfig(".env")
	cfg := kafka.Writer{
		Addr:                   kafka.TCP(cfgKafka.Host),
		Balancer:               kafka.CRC32Balancer{},
		AllowAutoTopicCreation: true,
	}
	return producerImpl{
		k: &cfg,
	}
}

func newKafkaWriter(topic string) *kafka.Writer {
	if kafkaWriters[topic] == nil {
		cfgKafka, _ := config.NewKafkaConfig(".env")
		kafkaWriters[topic] = &kafka.Writer{
			Addr:                   kafka.TCP(cfgKafka.Host),
			Topic:                  topic,
			Balancer:               &kafka.CRC32Balancer{},
			AllowAutoTopicCreation: true,
		}
	}

	return kafkaWriters[topic]
}

func (p producerImpl) PublishMessage(topic string, messages []string) {
	fmt.Println("This is using segmentio")
	syncMutex.Lock()
	kafkaWriter := newKafkaWriter(topic)
	var kafkaMsgs []kafka.Message
	for _, msg := range messages {
		kafkaMsg := p.buildMessage(msg)
		kafkaMsgs = append(kafkaMsgs, kafkaMsg)
	}
	const retries = 3
	for i := 0; i < retries; i++ {
		err := kafkaWriter.WriteMessages(context.Background(), kafkaMsgs...)
		if errors.Is(err, kafka.LeaderNotAvailable) || errors.Is(err, context.DeadlineExceeded) {
			time.Sleep(time.Millisecond * 250)
			continue
		}
		if err != nil {
			fmt.Printf("Error when write messages kafka: %v", err)
		} else {
			fmt.Printf("produced %d message for %s", len(kafkaMsgs), topic)
			break
		}
	}
	fmt.Println("Done")
	syncMutex.Unlock()
}

// SendMessage function for sending message to topic kafka
func (p producerImpl) SendMessage(topic string, messages []string) {
	fmt.Println("This is using segmentio")
	total := len(messages)
	chunk := 100
	for i := 0; i < len(messages); i += chunk {
		batch := messages[i:int(math.Min(float64(i+chunk), float64(total)))]
		fmt.Printf("produce %d of %d to %s", i+len(batch), total, topic)
		for _, msg := range batch {
			var err error
			const retries = 3
			for i := 0; i < retries; i++ {
				_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				defer cancel()
				// build message
				kafkaMsg := p.buildMessage(msg)
				err = p.k.WriteMessages(
					context.Background(),
					kafkaMsg,
				)
				if errors.Is(err, kafka.LeaderNotAvailable) || errors.Is(err, context.DeadlineExceeded) {
					time.Sleep(time.Millisecond * 250)
					continue
				}
				if err != nil {
					fmt.Printf("unexpected error %v", err)
				}
			}
			if err := p.k.Close(); err != nil {
				fmt.Printf("failed to close writer: %v", err)
			}
		}
	}
}

// buildMessage private function for building message to send to kafka
func (p producerImpl) buildMessage(messages string) kafka.Message {
	uuidV4, _ := uuid.NewV4()
	key := fmt.Sprintf("Key-%s", uuidV4.String())
	kafkaMsgs := kafka.Message{
		Key:   []byte(key),
		Value: []byte(messages),
	}

	return kafkaMsgs
}
