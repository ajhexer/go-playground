package service

import (
	"Banking/domain"
	"Banking/errors"
	"Banking/transfer"
)

type CustomerService interface{
	GetAllCustomer() ([]domain.Customer, *errors.AppError)
	GetCustomer(id string)(*transfer.CustomerResponse, *errors.AppError)
}

type DefaultCustomerService struct{
	repo domain.CustomerRepository
}

func (d DefaultCustomerService) GetAllCustomer() ([]transfer.CustomerResponse, *errors.AppError)  {
	customers, err := d.repo.FindAll()
	if err!=nil{
		return nil, err
	}
	res:=make([]transfer.CustomerResponse, 0)
	for _, customer := range customers{
		res = append(res, customer.ToDt())
	}
	return res, nil
}
func (d DefaultCustomerService) GetCustomer(id string )(*transfer.CustomerResponse, *errors.AppError){
	customer, err := d.repo.ById(id)
	if err!=nil{
		return nil, err
	}
	res := customer.ToDt()
	return &res, nil
}

func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService{
	return DefaultCustomerService{repo}
}
