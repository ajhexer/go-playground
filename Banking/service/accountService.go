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

	if r.TransactionType=="Withdrawal"{
		account, err := d.repo.FindById(r.AccountId)
		if err!=nil{
			return nil, err
		}
		if !account.CanWithdraw(r.Amount){
			return nil, errors.ValidationError("Credit is lower than amount")
		}
	}
	transaction := domain.NewTransaction(r.AccountId, r.TransactionType, r.Amount)
	t, err:=d.repo.SaveTransaction(transaction)
	if err!=nil{
		return nil, err
	}
	return t.ToDto(), nil
}

func (d DefaultAccountService) NewAccount(request transfer.AccountRequest) (*transfer.AccountResponse, *errors.AppError) {
	if err:=request.Validation(); err!=nil{
		return nil, err
	}
	account := domain.NewAccount(request.CustomerId, request.AccountType, request.Amount)
	t, err := d.repo.Save(account)
	if err!=nil{
		return nil, err
	}
	return t.ToDtResponse(), nil
}








