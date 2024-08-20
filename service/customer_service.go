package service

import "github.com/236sec/hexagonal/repository"

type customerService struct {
	custRepo repository.CustomerRepository
}

func NewCustomerService(r repository.CustomerRepository) CustomerService {
	return customerService{custRepo: r}
}

func (c customerService) GetAll() ([]CustomerResponse, error) {
	return nil, nil
}

func (c customerService) FindUser(id int) (*CustomerResponse, error) {
	return nil, nil
}
