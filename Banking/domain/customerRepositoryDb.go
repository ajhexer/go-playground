package domain

import (
	"Banking/logger"
	"database/sql"
)

type CustomerRepositoryDb struct{
	client *sql.DB
}

func (c CustomerRepositoryDb)FindAll() ([]Customer, error){

}

func (c CustomerRepositoryDb) ById()(*Customer, error){

}



func NewCustomerRepositoryDb() CustomerRepositoryDb{
	client, err := sql.Open("mysql", "user:password@tcp(localhost:5555)/banking")
	if err!=nil{
		logger.Log.Error("Unexpected sql err")

	}
}