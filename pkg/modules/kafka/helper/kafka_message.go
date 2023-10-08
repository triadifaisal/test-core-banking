package helper

import (
	"encoding/json"
	"time"

	"github.com/gofrs/uuid"
	gUUID "github.com/google/uuid"
)

type PropertyIDMessage struct {
	MessageID uuid.UUID
	ID        gUUID.UUID
	Message   string // for deposit case will contains C, for withdraw case will contains D
	Nominal   float64
	Timestamp int64
}

func BuildKafkaMessage(uuidTarget gUUID.UUID, trxCode string, nominal float64) []string {
	var messages []string
	messageID, _ := uuid.NewV4()
	propertyIDMessage := PropertyIDMessage{
		MessageID: messageID,
		ID:        uuidTarget,
		Message:   trxCode,
		Nominal:   nominal,
		Timestamp: time.Now().UnixMilli(),
	}

	bMessage, _ := json.Marshal(propertyIDMessage)

	messages = append(messages, string(bMessage))
	return messages
}

// FromKafka will convert message from kafka to PropertyIDMessage
func (p PropertyIDMessage) FromKafka(value []byte) PropertyIDMessage {
	var data PropertyIDMessage
	if err := json.Unmarshal(value, &data); err != nil {
		return data
	}
	return data
}
