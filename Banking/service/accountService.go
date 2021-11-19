package service

import (
	"Banking/domain"
	"Banking/errors"
	"Banking/transfer"
)

type AccountService interface{
	NewTransaction(r transfer.TransactionRequest) (*transfer.TransactionResponse, *errors.AppError)
	NewAccount(request transfer.AccountRequest)(*transfer.AccountResponse, *errors.AppError)
}

type DefaultAccountService struct{
	repo domain.AccountRepositoryDb
}


func (d DefaultAccountService) NewTransaction(r transfer.TransactionRequest) (*transfer.TransactionResponse, *errors.AppError) {
	panic("implement me")
}

func (d DefaultAccountService) NewAccount(request transfer.AccountRequest) (*transfer.AccountResponse, *errors.AppError) {
	panic("implement me")
}








