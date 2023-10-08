package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"core-banking/config"
	trx "core-banking/internal/transaction/service"
	"core-banking/pkg/dto/request"
	topicKafka "core-banking/pkg/modules/kafka"
	"core-banking/pkg/modules/kafka/helper"
	"core-banking/pkg/repository"

	"gorm.io/gorm"

	"github.com/labstack/gommon/log"
	"github.com/segmentio/kafka-go"
)

var (
	// kafka
	kafkaTopic string
)

func main() {
	flag.StringVar(&kafkaTopic, "topic", string(topicKafka.MutationIDs), "Kafka topic. Only one topic per worker.")

	flag.Parse()

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	cfgKafka, _ := config.NewKafkaConfig(".env")
	brokers := strings.Split(cfgKafka.Host, ",")

	// Section connect to DB
	// init config
	cfg, _ := config.NewConfig(".env")

	// init db
	pgxPool, _ := config.NewPgx(*cfg)

	// make a new reader that consumes from topic-A
	config := kafka.ReaderConfig{
		Brokers:         brokers,
		GroupID:         cfgKafka.GroupID,
		Topic:           kafkaTopic,
		MinBytes:        10e3,            // 10KB
		MaxBytes:        10e6,            // 10MB
		MaxWait:         1 * time.Second, // Maximum amount of time to wait for new data to come when fetching batches of messages from kafka.
		ReadLagInterval: -1,
	}

	reader := kafka.NewReader(config)
	defer reader.Close()

	for {
		var msg = helper.PropertyIDMessage{}
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Error(fmt.Sprintf("error while receiving message: %s", err.Error()))
			continue
		}

		value := m.Value

		msg = msg.FromKafka(value)

		process(context.Background(), pgxPool, msg)

		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s\n", m.Topic, m.Partition, m.Offset, string(value))
	}
}

func process(ctx context.Context, db *gorm.DB, kafkaMsg helper.PropertyIDMessage) {
	// init repo
	mutationRepo := repository.NewMutationRepository(db)

	//init service
	trxSvc := trx.NewTransactionService(mutationRepo)

	mutationReq := request.MutationRequest{
		UserUUID: kafkaMsg.ID,
		TrxCode:  kafkaMsg.Message,
		TrxTime:  time.Unix(0, kafkaMsg.Timestamp*int64(time.Millisecond)),
		Nominal:  kafkaMsg.Nominal,
	}

	isSuccess, err := trxSvc.Create(ctx, mutationReq)
	if !isSuccess {
		log.Error(fmt.Sprintf("error while create mutation: %s", err.Error()))
	}

	log.Infof("Finish process insert from kafka with User ID %v", mutationReq)
}
