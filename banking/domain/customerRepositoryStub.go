package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{
			Id: "1001",
			Name: "Shakti",
			City: "Blr",
			ZipCode: "560103",
			DateOfBirth: "01-08-1997",
			Status: "1",
		},
	}

	return CustomerRepositoryStub{customers}
}
