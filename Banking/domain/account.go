package domain

import (
	"Banking/errors"
	"Banking/transfer"
)

type Account struct{
	AccountId	string
	CustomerId	string
	OpeningDate	string
	AccountType	string
	Amount		float64
	Status		string
}

type AccountRepository interface{
	FindById(accountId string)(*Account, *errors.AppError)
	Save(account Account) *errors.AppError
}

func (a Account) ToDtResponse() *transfer.AccountResponse{
	return &transfer.AccountResponse{
		AccountId: a.AccountId,
	}
}

