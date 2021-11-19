package domain

import (
	"Banking/errors"
	"Banking/logger"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (a AccountRepositoryDb) FindById(accountId string)  (*Account, *errors.AppError){
	query := "SELECT account_id, customer_id, opening_date, account_type, amount from accounts where account_id = ?"
	var account Account
	err := a.client.Get(&account, query, accountId)
	if err != nil {
		logger.Log.Error("Error while getting account data from database"+err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}
	return &account, nil
}

func (a AccountRepositoryDb) Save(account Account) (*Account,*errors.AppError) {
	query :=  "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?)"
	result, err := a.client.Exec(query, account.CustomerId, account.OpeningDate, account.AccountType, account.Amount, account.Status)
	if err!=nil{
		logger.Log.Error("Error while inserting account into database: "+err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}
	id, err:=result.LastInsertId()
	if err!=nil{
		logger.Log.Error("Error while getting account id from database: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}
	account.AccountId = strconv.FormatInt(id, 10)
	return &account, nil
}

func (a AccountRepositoryDb) SaveTransaction(t Transaction) (*Transaction, *errors.AppError) {
	tx, err := a.client.Begin()
	if err != nil {
		logger.Log.Error("Error while starting a new transaction for bank account transaction: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}

	result, _ := tx.Exec(`INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) 
											values (?, ?, ?, ?)`, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)

	if t.TransactionType == "Withdrawal" {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount - ? where account_id = ?`, t.Amount, t.AccountId)
	} else {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount + ? where account_id = ?`, t.Amount, t.AccountId)
	}

	if err != nil {
		tx.Rollback()
		logger.Log.Error("Error while saving transaction: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Log.Error("Error while adding transaction to account: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}
	transactionId, err := result.LastInsertId()
	if err != nil {
		logger.Log.Error("Error while getting transaction id: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}

	account, appErr := a.FindById(t.AccountId)
	if appErr != nil {
		return nil, appErr
	}
	t.TransactionId = strconv.FormatInt(transactionId, 10)

	t.Amount = account.Amount
	return &t, nil
}


func NewAccountRepositoryDb(client *sqlx.DB) AccountRepositoryDb{
	return AccountRepositoryDb{client: client}
}




