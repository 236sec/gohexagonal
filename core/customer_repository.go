package core

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	FindById() (*Customer, error)
	Create(CreateCustomerDto) error
}