package service

import (
	"context"
	"core-banking/pkg/dto/request"
	"core-banking/pkg/modules/kafka"
	"core-banking/pkg/modules/kafka/helper"
)

func (svc *UserService) Withdraw(ctx context.Context, req request.WithdrawRequest) (float64, error) {
	data, err := svc.repo.GetByAccountNumber(ctx, req.AccountNumber)
	if err != nil {
		return float64(0), err
	}

	if errSubtract := data.SubtractBalance(req.Nominal); errSubtract != nil {
		return float64(0), errSubtract
	}

	if err := svc.repo.UpdateBalance(ctx, req.AccountNumber, data.Balance); err != nil {
		return float64(0), err
	}

	// Publish message to kafka
	producer := kafka.NewProducer()
	producer.PublishMessage(string(kafka.MutationIDs), helper.BuildKafkaMessage(data.UUID, "D", req.Nominal))

	return data.Balance, nil
}
