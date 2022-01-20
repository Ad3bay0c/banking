package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (c CustomerRepositoryStub) FindAll() ([]Customer, error)  {
	return c.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{
			ID:      "1",
			Name:    "John Doe",
			City:    "New York",
			Zipcode: "10001",
			Dob:     "01/01/1990",
			Status:  "active",
		},
		{
			ID:      "2",
			Name:    "Jane Doe",
			City:    "New York",
			Zipcode: "10001",
			Dob:     "01/01/1990",
			Status:  "active",
		},
	}

	return CustomerRepositoryStub{
		customers: customers,
	}
}
