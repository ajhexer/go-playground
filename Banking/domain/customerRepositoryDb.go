package domain

import (
	"Banking/errors"
	"Banking/logger"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct{
	client *sqlx.DB
}

func (c CustomerRepositoryDb)FindAll() ([]Customer, *errors.AppError){
	query := "select customer_id, name, date_of_birth, city, zipcode, status from customers"
	customers := make([]Customer, 0)
	err := c.client.Select(&customers, query)
	if err!=nil{
		logger.Log.Error("Error while scanning database rows")
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}
	return customers, nil
}

func (c CustomerRepositoryDb) ById(id string)(*Customer, *errors.AppError){
	query := "select customer_id, name, data_of_birth, city, zipcode, status from customers where customer_id=?"
	var customer Customer
	err := c.client.Get(&customer, query, id)
	if err!=nil{
		logger.Log.Error("Error while querying customer by id")
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}
	return &customer, nil
}



func NewCustomerRepositoryDb(client *sqlx.DB) CustomerRepositoryDb{
	return CustomerRepositoryDb{
		client: client,
	}
}