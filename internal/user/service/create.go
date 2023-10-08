package service

import (
	"context"
	"core-banking/pkg/dto/request"
	"core-banking/pkg/model"
	"core-banking/pkg/reference"
	"crypto/rand"
	"database/sql"
	"errors"
	"fmt"
)

// define the given charset for no rekening
const (
	charSet             = "1234567890"
	accountNumberLength = 10
	maxRetry            = 3
)

func (svc *UserService) Create(ctx context.Context, req request.CreateRequest) (string, error) {
	data, err := svc.repo.CheckNikAndPhoneNumberExist(ctx, req.NIK, req.PhoneNumber)
	if err != nil {
		return "", errors.New("[CheckNikAndPhoneNumberExist] error: " + err.Error())
	}

	if data != nil {
		return "", reference.ErrUniqueHpNik
	}

	accountNumber := ""
	i := 0
	for {
		// maximum retry generate account number
		if i > maxRetry {
			return "", errors.New("system error when generate account number")
		}

		// Generating account number
		accountNumber, err = generateAccountNumber(req.NIK, req.PhoneNumber)
		if err != nil {
			return "", reference.ErrGenerateAccountNumber
		}

		// CHeck account number is exist or not on DB
		accountNumberExist, err := svc.repo.CheckAccountNumberExist(ctx, accountNumber)
		if err == nil && *accountNumberExist {
			i++
			continue
		} else if err != nil {
			return "", err
		}

		err = svc.repo.Insert(ctx, &model.User{
			Name: req.Name,
			NIK:  req.NIK,
			Phonenumber: sql.NullString{
				String: req.PhoneNumber,
				Valid:  true,
			},
			AccountNumber: accountNumber,
			Balance:       float64(0),
		})

		break
	}
	return accountNumber, nil
}

func generateAccountNumber(nik, phoneNumber string) (string, error) {
	bytes := make([]byte, accountNumberLength)
	copy(bytes[:], fmt.Sprintf("%s%s", nik, phoneNumber))
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = charSet[b%byte(len(charSet))]
	}
	return string(bytes), nil
}
