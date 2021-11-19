package transfer

import "Banking/errors"

type AccountRequest struct {
	CustomerId	string
	Amount		float64
	AccountType	string
}

func (a AccountRequest) Validation() *errors.AppError{
	if a.AccountType!="saving" && a.AccountType!="salary"{
		return errors.
	}
}