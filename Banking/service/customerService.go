package service

import "Banking/domain"

type CustomerService interface{
	GetAllCustomer() ([]domain.Customer, error)
}

type DefaultCustomerService struct{
	repo domain.CustomerRepository
}

func (d DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error)  {
	return d.repo.FindAll()
}

func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService{
	return DefaultCustomerService{repo}
}
