package service

import "github.com/236sec/hexagonal/repository"

type customerService struct {
	custRepo repository.CustomerRepository
}

func NewCustomerService(r repository.CustomerRepository) CustomerService {
	return &customerService{custRepo: r}
}

func (s customerService) CreateUser(customer CustomerResponse) error {
	if err := s.custRepo.Create(customer); err != nil {
		return err
	}
	return nil
}

func (s customerService) GetAll() ([]CustomerResponse, error) {
	customers, err := s.custRepo.GetAll()
	if err != nil {
		return nil, err
	}
	custResponses := []CustomerResponse{}
	for _, customer := range customers {
		custResponse := CustomerResponse{
			CustomerID: int(customer.CustomerID),
			Name:       customer.Name,
			Email:      customer.Email,
			Amount:     customer.Amount,
		}
		custResponses = append(custResponses, custResponse)
	}
	return custResponses, nil
}

func (s customerService) FindUser(id int) (*CustomerResponse, error) {
	customer, err := s.custRepo.GetById(id)
	if err != nil {
		return nil, err
	}
	custResponse := CustomerResponse{
		CustomerID: int(customer.CustomerID),
		Name:       customer.Name,
		Email:      customer.Email,
		Amount:     customer.Amount,
	}
	return &custResponse, nil
}
