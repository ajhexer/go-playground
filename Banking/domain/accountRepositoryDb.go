package domain

import (
	"Banking/errors"
	"Banking/logger"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type AccountRepositoryDb struct {
	client sqlx.DB
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



