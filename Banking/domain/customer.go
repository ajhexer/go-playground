package domain

import (
	"Banking/errors"
	"Banking/transfer"
)

type Customer struct {
	Id			string
	Name		string
	City		string
	Zipcode		string
	DateOfBirth	string
	Status 		string

}
type CustomerRepository interface{
	FindAll() ([]Customer, *errors.AppError)
	ById(id string) (*Customer, *errors.AppError)
}

func (c Customer) ToDt() transfer.CustomerResponse{
	res := transfer.CustomerResponse{
		c.Id,
		c.Name,
		c.City,
		c.Zipcode,
		c.DateOfBirth,
		c.transStatus(),
	}
	return res
}
func (c Customer) transStatus() string{
	if c.Status=="1"{
		return "active"
	}
	return "inactive"
}
