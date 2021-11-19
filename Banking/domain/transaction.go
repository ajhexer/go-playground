package domain

import (
	"Banking/transfer"
	"time"
)

type Transaction struct {
	TransactionId   string  `db:"transaction_id"`
	AccountId       string  `db:"account_id"`
	Amount          float64 `db:"amount"`
	TransactionType string  `db:"transaction_type"`
	TransactionDate string  `db:"transaction_date"`
}



func (t Transaction) ToDto() *transfer.TransactionResponse {
	return &transfer.TransactionResponse{
		TransactionId:   t.TransactionId,
		AccountId:       t.AccountId,
		Amount:          t.Amount,
		TransactionType: t.TransactionType,
		TransactionDate: t.TransactionDate,
	}
}

func NewTransaction(accountId string, transactionType string, amount float64) Transaction{
	return Transaction{
		AccountId: accountId,
		Amount: amount,
		TransactionType: transactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}
}

