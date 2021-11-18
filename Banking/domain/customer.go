package domain

import "Banking/errors"

type Customer struct {
	Id			string
	Name		string
	City		string
	Zipcode		string
	DateOfBirth	string
	status 		string

}
type CustomerRepository interface{
	FindAll() ([]Customer, *errors.AppError)
	ById(id int) (*Customer, *errors.AppError)
}

