package core

type CustomerService interface {
	GetAll() ([]Customer, error)
	FindUser() (*Customer, error)
	CreateUser(CreateCustomerDto) error
}

type customerImplement struct {
	repo CustomerRepository
}

func NewCustomerService(repo CustomerRepository) CustomerService {
	return &customerImplement{repo: repo}
}

func (s *customerImplement) GetAll() ([]Customer, error) {
	return nil, nil
}

func (s *customerImplement) FindUser() (*Customer, error) {
	return nil, nil
}

func (s *customerImplement) CreateUser(customer CreateCustomerDto) error {
	err := s.repo.Create(customer)
	if err != nil {
		return err
	}
	return nil
}