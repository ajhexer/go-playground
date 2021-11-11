package domain

type CustomerRepositoryStub struct{
	customers [] Customer
}

func (c CustomerRepositoryStub) FindAll() ([]Customer, error)  {
	return c.customers, nil
}
func NewCustomerRepositoryStub() CustomerRepositoryStub{
	customers := []Customer{
		{"1001", "Arash", "Tehran", "33314", "2001-02-01", "1"},
		{"1002", "Ahmad", "New york", "10001", "1980-07-08", "1"},
	}
	return CustomerRepositoryStub{customers: customers}
}
