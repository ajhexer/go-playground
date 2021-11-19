package domain

import (
	"Banking/errors"
	"Banking/transfer"
	"time"
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
	Save(account Account) (*Account, *errors.AppError)
	SaveTransaction(transaction Transaction) (*Transaction, *errors.AppError)
}

func (a Account) ToDtResponse() *transfer.AccountResponse{
	return &transfer.AccountResponse{
		AccountId: a.AccountId,
	}
}

func NewAccount(customerId string, accountType string, amount float64) Account{
	return Account{
		CustomerId: customerId,
		AccountType: accountType,
		Amount: amount,
		Status: "1",
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
	}
}

func (a Account) CanWithdraw(amount float64) bool {
	return amount >= a.Amount
}

