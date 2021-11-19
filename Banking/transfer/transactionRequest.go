package transfer


type TransactionRequest struct{
	Amount			float64	`json:"amount"`
	TransactionType	string	`json:"transaction_type"`
	AccountId 		string	`json:"account_id"`
}
func (t TransactionRequest) Validate(accountId string) bool{
	return accountId==t.AccountId
}






