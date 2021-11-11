package domain

type Customer struct {
	Id			string
	Name		string
	City		string
	Zipcode		string
	DateOfBirth	string
	status 		string

}
type CustomerRepository interface{
	FindAll() ([]Customer, error)
}

